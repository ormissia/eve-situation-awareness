package com.ormissia.zkill.analyzer

import com.ormissia.zkill.sink.MySQLSink
import com.ormissia.zkill.source.KafkaSource
import com.ormissia.zkill.transformation.KafkaLineToZKillInfo
import com.ormissia.zkill.utils.{ESAConst, SolarSystemSink, ZKillInfo}
import org.apache.flink.api.common.functions.AggregateFunction
import org.apache.flink.streaming.api.CheckpointingMode
import org.apache.flink.streaming.api.environment.CheckpointConfig
import org.apache.flink.streaming.api.scala.function.WindowFunction
import org.apache.flink.streaming.api.windowing.assigners.SlidingEventTimeWindows
import org.apache.flink.streaming.api.windowing.time.Time
import org.apache.flink.streaming.api.windowing.windows.TimeWindow
import org.apache.flink.util.Collector
import org.apache.flink.streaming.api.scala._
import org.slf4j.LoggerFactory

import java.text.SimpleDateFormat
import java.util.TimeZone


object SolarSystemKillAnalyzer {
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

    // 0         1      2    3    4   5           6        7      8         9
    // timeStamp killId hash solo npc solarSystem iskValue victim attackers labels...
    val zkillInfoStream = KafkaSource.GetZKillInfoStream(env, this.getClass.getName)

    val result = zkillInfoStream
      .map(new KafkaLineToZKillInfo).uid("ConvertToZkillInfo").setParallelism(3)
      .keyBy(_.solarSystem)
      .window(
        SlidingEventTimeWindows.of(
          Time.hours(1),
          Time.minutes(10)
        ))
      // 在这里允许延迟数据重新计算会导致结果错误，延迟数据全部使用侧输出流处理
      //.allowedLateness(Time.hours(2))
      // 延迟严重的数据到侧输出流
      .sideOutputLateData(KafkaSource.LATE_DATE_TAG_ZKILL_INFO)
      .aggregate(
        new AggregateFunction[ZKillInfo, SolarSystemSink, SolarSystemSink] {
          override def createAccumulator(): SolarSystemSink = SolarSystemSink(0, 0, 0, "")

          override def add(value: ZKillInfo, accumulator: SolarSystemSink): SolarSystemSink = {
            SolarSystemSink(
              value.solarSystem,
              accumulator.killQuantity + 1,
              accumulator.killValue + value.totalValue,
              "",
            )
          }

          override def getResult(accumulator: SolarSystemSink): SolarSystemSink = accumulator

          override def merge(a: SolarSystemSink, b: SolarSystemSink): SolarSystemSink = {
            SolarSystemSink(
              a.solarSystemId,
              a.killQuantity + b.killQuantity,
              a.killValue + b.killValue,
              "",
            )
          }
        },
        new WindowFunction[SolarSystemSink, SolarSystemSink, Int, TimeWindow] {
          override def apply(key: Int, window: TimeWindow, input: Iterable[SolarSystemSink],
                             out: Collector[SolarSystemSink]): Unit = {
            val dateFormat = new SimpleDateFormat(ESAConst.DATE_FORMAT_yyyyMMddHH)
            dateFormat.setTimeZone(TimeZone.getTimeZone("UTC"))
            var dt = dateFormat.format(window.getEnd)
            if (window.getEnd % 3600000 == 0) {
              dt = dateFormat.format(window.getStart)
            }

            val result = SolarSystemSink(input.head.solarSystemId, input.head.killQuantity, input.head.killValue, dt)
            out.collect(result)

            val dateFormat1 = new SimpleDateFormat(ESAConst.DATE_FORMAT_yyyyMMddHHmmss)
            dateFormat1.setTimeZone(TimeZone.getTimeZone("UTC"))
            val start = dateFormat1.format(window.getStart)
            val end = dateFormat1.format(window.getEnd)
            LOG.info(s"start: ${start}\tend: ${end}\t${result}")
          }
        }
      ).setParallelism(3)

    result.addSink(new MySQLSink[SolarSystemSink](classOf[SolarSystemSink])).setParallelism(3)

    // TODO 处理延迟严重的数据
    //result.getSideOutput(KafkaSource.LATE_DATE_TAG_ZKILL_INFO).print("late")

    env.execute(this.getClass.getName)
  }
}
