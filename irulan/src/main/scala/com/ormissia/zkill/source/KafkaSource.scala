package com.ormissia.zkill.source

import org.apache.flink.api.common.serialization.SimpleStringSchema
import org.apache.flink.streaming.connectors.kafka.{FlinkKafkaConsumer, FlinkKafkaConsumerBase}
import org.apache.kafka.common.serialization.StringSerializer

import java.util.Properties

object KafkaSource {
  def GetKafkaSource(): FlinkKafkaConsumerBase[String] = {
    // TODO 配置传入方式需要修改
    val properties = new Properties()
    properties.setProperty("bootstrap.servers", "192.168.13.107:9092,192.168.13.108:9092,192.168.13.109:9092")
    properties.setProperty("group.id", "KillValue-test")
    properties.setProperty("key.deserializer", classOf[StringSerializer].getName)
    properties.setProperty("value.deserializer", classOf[StringSerializer].getName)

    val consumer = new FlinkKafkaConsumer[String]("zkill", new SimpleStringSchema(), properties)
    // TODO 设置消费方式
    consumer.setStartFromEarliest()
    //consumer.setStartFromGroupOffsets()
  }
}
