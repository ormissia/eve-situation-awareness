package com.ormissia.zkill.analyzer

import com.ormissia.zkill.source.KafkaSource
import com.ormissia.zkill.transformation.KafkaLineToZKillInfo
import com.ormissia.zkill.utils.{Attacker, CharacterSink, ZKillInfo}
import org.apache.flink.api.common.functions.AggregateFunction
import org.apache.flink.streaming.api.scala._
import org.apache.flink.streaming.api.scala.function.ProcessAllWindowFunction
import org.apache.flink.streaming.api.windowing.assigners.SlidingEventTimeWindows
import org.apache.flink.streaming.api.windowing.time.Time
import org.apache.flink.streaming.api.windowing.windows.TimeWindow
import org.apache.flink.util.Collector

import scala.collection.mutable.ListBuffer

object CharacterKillAnalyzer {
  def main(args: Array[String]): Unit = {
    val env = StreamExecutionEnvironment.getExecutionEnvironment

    val zkillInfoStream = KafkaSource.GetZKillInfoStream(env, this.getClass.getName)
    zkillInfoStream
      .map(new KafkaLineToZKillInfo).uid("ConvertToZkillInfo")
      .flatMap(zKillInfo => {
        val res = new ListBuffer[Attacker]
        zKillInfo.attackers.foreach(attacker => {
          res += Attacker(
            characterId = attacker.characterId,
            killQuantity = 1,
            shipTypeId = attacker.shipTypeId,
            solarSystem = zKillInfo.solarSystem,
            killValue = zKillInfo.totalValue,
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
      .sideOutputLateData(KafkaSource.LATE_DATE_TAG_ATTACKER)
      //.process(new ProcessAllWindowFunction[Attacker,CharacterSink,TimeWindow] {
      //  override def process(context: Context, elements: Iterable[Attacker], out: Collector[CharacterSink]): Unit = {
      //    out.collect()
      //  }
      //})
      //.aggregate(
      //  new AggregateFunction[Attacker, CharacterSink, CharacterSink] {
      //    override def createAccumulator(): CharacterSink = CharacterSink()
      //
      //    override def add(value: Attacker, accumulator: CharacterSink): CharacterSink = ???
      //
      //    override def getResult(accumulator: CharacterSink): CharacterSink = accumulator
      //
      //    override def merge(a: CharacterSink, b: CharacterSink): CharacterSink = ???
      //  }
      //)

    env.execute(this.getClass.getName)
  }
}
