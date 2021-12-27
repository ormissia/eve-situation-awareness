package com.ormissia.zkill.analyzer

import com.ormissia.zkill.sink.MySQLSink
import com.ormissia.zkill.source.KafkaSource
import com.ormissia.zkill.transformation.KafkaLineToZKillInfo
import com.ormissia.zkill.utils.{ESAConst, SolarSystemSink, ZKillInfo}
import org.apache.flink.api.common.eventtime.{SerializableTimestampAssigner, WatermarkStrategy}
import org.apache.flink.api.common.functions.AggregateFunction
import org.apache.flink.streaming.api.scala.function.WindowFunction
import org.apache.flink.streaming.api.windowing.assigners.SlidingEventTimeWindows
import org.apache.flink.streaming.api.windowing.time.Time
import org.apache.flink.streaming.api.windowing.windows.TimeWindow
import org.apache.flink.util.Collector
import org.apache.flink.streaming.api.scala._

import java.text.SimpleDateFormat
import java.time.Duration
import java.util.TimeZone


object KillValue {
  def main(args: Array[String]): Unit = {

    val env = StreamExecutionEnvironment.getExecutionEnvironment

    // TODO 开启检查点
    //设置Checkpoint间隔
    //env.enableCheckpointing(1000)
    //Checkpoint之间的最小时间间隔
    //env.getCheckpointConfig.setMinPauseBetweenCheckpoints(500)

    val kafkaConsumer = KafkaSource.GetKafkaSource()
    val zkillKafkaStream = env.addSource(kafkaConsumer).setParallelism(3)
      //val zkillKafkaStream = env.socketTextStream("127.0.0.1", 1234)
      .assignTimestampsAndWatermarks(
        WatermarkStrategy.forBoundedOutOfOrderness[String](Duration.ofHours(1))
          //.withIdleness(Duration.ofMinutes(10))
          .withTimestampAssigner(new SerializableTimestampAssigner[String] {
            override def extractTimestamp(element: String, recordTimestamp: Long): Long = {
              element.split(" ").head.toLong
            }
          })
      )
    //val zkillKafkaStream: DataStream[String] = env.socketTextStream("127.0.0.1", 1234)

    // 0         1      2    3    4   5           6        7      8         9
    // timeStamp killId hash solo npc solarSystem iskValue victim attackers labels...
    zkillKafkaStream
      // TODO 增加侧输出流保存不合法的数据
      // 初步过滤不合法的数据流
      .filter(_.split(" ").length >= 8)
      // 将kafka的每一行数据转换成ZKillInfo
      .map(new KafkaLineToZKillInfo).uid("ConvertToZKillInfo")
      .keyBy(_.solarSystem)
      .window(
        SlidingEventTimeWindows.of(
          Time.hours(1),
          Time.minutes(10)
        ))
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
          override def apply(key: Int, window: TimeWindow, input: Iterable[SolarSystemSink], out: Collector[SolarSystemSink]): Unit = {
            val dateFormat = new SimpleDateFormat(ESAConst.DATE_FORMAT_yyyyMMddHH)
            dateFormat.setTimeZone(TimeZone.getTimeZone("UTC"))
            var dt = dateFormat.format(window.getEnd)
            if (window.getEnd % 3600000 == 0) {
              dt = dateFormat.format(window.getStart)
            }

            val result = SolarSystemSink(input.head.solarSystemId, input.head.killQuantity, input.head.killValue, dt)
            out.collect(result)

            // TODO log
            val dateFormat1 = new SimpleDateFormat(ESAConst.DATE_FORMAT_yyyyMMddHHmmss)
            dateFormat1.setTimeZone(TimeZone.getTimeZone("UTC"))
            val start = dateFormat1.format(window.getStart)
            val end = dateFormat1.format(window.getEnd)
            println(s"start: ${start}\tend: ${end}\t${result}")
          }
        }
      )
      .addSink(new MySQLSink[SolarSystemSink](classOf[SolarSystemSink]))

    env.execute("KillValue")
  }
}
