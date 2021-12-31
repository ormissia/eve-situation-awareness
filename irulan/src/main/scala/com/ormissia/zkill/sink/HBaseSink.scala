package com.ormissia.zkill.sink

import org.apache.flink.configuration.Configuration
import org.apache.flink.streaming.api.functions.sink.{RichSinkFunction, SinkFunction}

class HBaseSink extends RichSinkFunction[String]{
  override def open(parameters: Configuration): Unit = {

  }

  override def close(): Unit = {

  }

  override def invoke(value: String, context: SinkFunction.Context): Unit = {

  }
}
