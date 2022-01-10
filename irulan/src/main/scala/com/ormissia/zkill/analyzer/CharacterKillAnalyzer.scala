package com.ormissia.zkill.analyzer

import com.ormissia.zkill.sink.MySQLSink
import com.ormissia.zkill.source.KafkaSource
import com.ormissia.zkill.transformation.KafkaLineToZKillInfo
import com.ormissia.zkill.utils.{Attacker, CharacterSink, ESAConst}
import org.apache.flink.api.common.functions.AggregateFunction
import org.apache.flink.api.common.state.MapState
import org.apache.flink.streaming.api.CheckpointingMode
import org.apache.flink.streaming.api.environment.CheckpointConfig
import org.apache.flink.streaming.api.scala._
import org.apache.flink.streaming.api.scala.function.{ProcessAllWindowFunction, WindowFunction}
import org.apache.flink.streaming.api.windowing.assigners.SlidingEventTimeWindows
import org.apache.flink.streaming.api.windowing.time.Time
import org.apache.flink.streaming.api.windowing.windows.TimeWindow
import org.apache.flink.util.Collector
import org.slf4j.LoggerFactory

import java.text.SimpleDateFormat
import java.util.TimeZone
import scala.collection.mutable.ListBuffer

object CharacterKillAnalyzer {
  private val LOG = LoggerFactory.getLogger(this.getClass)

  def main(args: Array[String]): Unit = {
    val env = StreamExecutionEnvironment.getExecutionEnvironment


    env.getCheckpointConfig.setCheckpointStorage("hdfs://bigdata/flink/checkpoint/" + this.getClass.getName)
    // 设置Checkpoint间隔
    env.enableCheckpointing(60 * 1000)
    //Checkpoint之间的最小时间间隔
    env.getCheckpointConfig.setMinPauseBetweenCheckpoints(1000)
    env.getCheckpointConfig.setCheckpointingMode(CheckpointingMode.EXACTLY_ONCE)
    // RETAIN_ON_CANCELLATION flink任务取消后，checkpoint数据会被保留
    env.getCheckpointConfig.enableExternalizedCheckpoints(
      CheckpointConfig.ExternalizedCheckpointCleanup.RETAIN_ON_CANCELLATION)

    val zkillInfoStream = KafkaSource.GetZKillInfoStream(env, this.getClass.getName)
    val result = zkillInfoStream
      .map(new KafkaLineToZKillInfo).uid("ConvertToZkillInfo").setParallelism(3)
      .flatMap(zKillInfo => {
        val res = new ListBuffer[Attacker]
        val labelsList = zKillInfo.labels.map((_, 1)).toList
        zKillInfo.attackers.foreach(attacker => {
          res += Attacker(
            characterId = attacker.characterId,
            killQuantity = 1,
            shipTypeId = attacker.shipTypeId,
            solarSystem = zKillInfo.solarSystem,
            killValue = zKillInfo.totalValue,
            finalBlow = attacker.finalBlow,
            labels = labelsList,
          )
        })
        res
      })
      .keyBy(_.characterId)
      .window(
        SlidingEventTimeWindows.of(
          Time.days(1),
          Time.minutes(10)
        ))
      // 延迟严重的数据到侧输出流
      // TODO
      .sideOutputLateData(KafkaSource.LATE_DATE_TAG_ATTACKER)
      .aggregate(
        new AggregateFunction[Attacker, CharacterSink, CharacterSink] {

          override def createAccumulator(): CharacterSink = CharacterSink(0, 0, 0, 0, "", List(), List(), List())

          override def add(value: Attacker, accumulator: CharacterSink): CharacterSink = {

            //println(s"valueLabels>>>>>${value.labels}\taccumulatorLabels>>>>>${accumulator.labels}\taddLabels>>>>>${value.labels+ accumulator.labels}")
            CharacterSink(
              value.characterId,
              accumulator.finalShoot + (if (value.finalBlow) 1 else 0),
              accumulator.killQuantity + 1,
              accumulator.killValue + value.killValue,
              "",
              value.labels ++ accumulator.labels,
              // TODO 使用map优化
              accumulator.shipTypes :+ (value.shipTypeId, 1),
              accumulator.solarSystems :+ (value.solarSystem, 1),
            )
          }

          override def getResult(accumulator: CharacterSink): CharacterSink = accumulator

          override def merge(a: CharacterSink, b: CharacterSink): CharacterSink = {
            CharacterSink(
              a.characterId,
              a.finalShoot + b.finalShoot,
              a.killQuantity + b.killQuantity,
              a.killValue + b.killValue,
              "",
              a.labels ++ b.labels,
              a.shipTypes ++ b.shipTypes,
              a.solarSystems ++ b.solarSystems,
            )
          }
        },
        new WindowFunction[CharacterSink, CharacterSink, Int, TimeWindow] {
          override def apply(key: Int, window: TimeWindow, input: Iterable[CharacterSink],
                             out: Collector[CharacterSink]): Unit = {
            val dateFormat = new SimpleDateFormat(ESAConst.DATE_FORMAT_yyyyMMddHH)
            dateFormat.setTimeZone(TimeZone.getTimeZone("UTC"))
            var dt = dateFormat.format(window.getEnd)
            if (window.getEnd % 3600000 == 0) {
              dt = dateFormat.format(window.getStart)
            }

            val result = CharacterSink(
              input.head.characterId,
              input.head.finalShoot,
              input.head.killQuantity,
              input.head.killValue,
              dt,
              input.head.labels,
              input.head.shipTypes,
              input.head.solarSystems)
            out.collect(result)

            val dateFormat1 = new SimpleDateFormat(ESAConst.DATE_FORMAT_yyyyMMddHHmmss)
            dateFormat1.setTimeZone(TimeZone.getTimeZone("UTC"))
            val start = dateFormat1.format(window.getStart)
            val end = dateFormat1.format(window.getEnd)
            LOG.info(s"start: ${start}\tend: ${end}\t${result}")
          }
        }
      )
      .map(character => {
        val labels = character.labels.groupBy(_._1).mapValues(_.size).toList.sortBy(_._2).reverse
        val shipTypes = character.shipTypes.groupBy(_._1).mapValues(_.size).toList.sortBy(_._2).reverse
        val solarSystems = character.solarSystems.groupBy(_._1).mapValues(_.size).toList.sortBy(_._2).reverse
        CharacterSink(
          character.characterId,
          character.finalShoot,
          character.killQuantity,
          character.killValue,
          character.dt,
          labels,
          shipTypes,
          solarSystems,
        )
      }).setParallelism(3)

    result.addSink(new MySQLSink[CharacterSink](classOf[CharacterSink])).setParallelism(3)

    env.execute(this.getClass.getName)
  }
}
