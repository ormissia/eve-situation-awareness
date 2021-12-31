package com.ormissia.zkill.source

import com.ormissia.zkill.utils.{Attacker, ZKillInfo}
import org.apache.flink.api.common.eventtime.{SerializableTimestampAssigner, WatermarkStrategy}
import org.apache.flink.api.common.serialization.SimpleStringSchema
import org.apache.flink.streaming.api.scala.{StreamExecutionEnvironment, _}
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer
import org.apache.kafka.common.serialization.StringSerializer

import java.time.Duration
import java.util.Properties

object KafkaSource {

  val LATE_DATE_TAG_ZKILL_INFO = new OutputTag[ZKillInfo]("late")
  val LATE_DATE_TAG_ATTACKER = new OutputTag[Attacker]("late")

  def GetZKillInfoStream(env: StreamExecutionEnvironment, groupName: String): DataStream[String] = {
    // TODO 配置传入方式需要修改
    val properties = new Properties()
    properties.setProperty("bootstrap.servers", "192.168.13.107:9092,192.168.13.108:9092,192.168.13.109:9092")
    properties.setProperty("group.id", groupName)
    properties.setProperty("key.deserializer", classOf[StringSerializer].getName)
    properties.setProperty("value.deserializer", classOf[StringSerializer].getName)

    val consumer = new FlinkKafkaConsumer[String]("zkill", new SimpleStringSchema(), properties)
    //consumer.setStartFromEarliest()
    consumer.setStartFromGroupOffsets()


    val zkillKafkaStream = env.addSource(consumer).setParallelism(3)
      //val zkillKafkaStream = env.socketTextStream("127.0.0.1", 1234)
      .assignTimestampsAndWatermarks(
        WatermarkStrategy.forBoundedOutOfOrderness[String](Duration.ofHours(1))
          .withTimestampAssigner(new SerializableTimestampAssigner[String] {
            override def extractTimestamp(element: String, recordTimestamp: Long): Long = {
              element.split(" ").head.toLong
            }
          })
      )

    zkillKafkaStream.name("Kafka Source line string")
      .filter(_.split(" ").length >= 8).setParallelism(3)
  }
}
