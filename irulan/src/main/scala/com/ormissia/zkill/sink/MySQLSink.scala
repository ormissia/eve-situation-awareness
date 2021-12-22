package com.ormissia.zkill.sink

import com.ormissia.zkill.utils.SolarSystemSink
import org.apache.flink.configuration.Configuration
import org.apache.flink.streaming.api.functions.sink.{RichSinkFunction, SinkFunction}

import java.sql.{Connection, DriverManager, PreparedStatement}

class MySQLSink[T](classType: Class[_ <: T]) extends RichSinkFunction[T] {
  var conn: Connection = _
  var pst: PreparedStatement = _

  override def invoke(value: T, context: SinkFunction.Context): Unit = {
    // TODO 判断方法需要修改
    if (classType.getName.equals(classOf[SolarSystemSink].getName)) {
      val info = value.asInstanceOf[SolarSystemSink]

      pst.setString(1, info.dt)
      pst.setInt(2, info.solarSystemId)
      pst.setInt(3, info.killQuantity)
      pst.setDouble(4, info.killValue)
      pst.setInt(5, info.killQuantity)
      pst.setDouble(6, info.killValue)

      pst.executeUpdate()
    }
  }

  override def open(parameters: Configuration): Unit = {
    conn = DriverManager.getConnection("jdbc:mysql://localhost/eve_situation_awareness", "root", "123")
    // TODO 判断方法需要修改
    if (classType.getName.equals(classOf[SolarSystemSink].getName)) {
      pst = conn.prepareStatement("insert into solar_system_kill_statistical (dt, solar_system_id, kill_quantity, kill_value, create_time) values (?, ?, ?, ?, now()) on duplicate key update kill_quantity = ?, kill_value= ?")
    }
  }

  override def close(): Unit = {
    pst.close()
    conn.close()
  }
}
