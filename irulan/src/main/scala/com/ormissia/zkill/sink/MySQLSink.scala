package com.ormissia.zkill.sink

import com.ormissia.zkill.utils.{CharacterSink, SolarSystemSink}
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
    } else if (classType.getName.equals(classOf[CharacterSink].getName)) {
      val info = value.asInstanceOf[CharacterSink]

      pst.setString(1, info.dt)
      pst.setInt(2, info.characterId)
      pst.setInt(3, info.finalShoot)
      pst.setInt(4, info.killQuantity)
      pst.setDouble(5, info.killValue)
      pst.setString(6, info.labels)
      pst.setString(7, info.shipTypes)
      pst.setString(8, info.solarSystems)

      pst.setInt(9, info.finalShoot)
      pst.setInt(10, info.killQuantity)
      pst.setDouble(11, info.killValue)
      pst.setString(12, info.labels)
      pst.setString(13, info.shipTypes)
      pst.setString(14, info.solarSystems)

      pst.executeUpdate()
    }
  }

  override def open(parameters: Configuration): Unit = {
    //conn = DriverManager.getConnection("jdbc:mysql://192.168.13.100:3306/eve_situation_awareness", "root", "123456")
    conn = DriverManager.getConnection("jdbc:mysql://127.0.0.1:3306/eve_situation_awareness", "root", "123")
    if (classType.getName.equals(classOf[SolarSystemSink].getName)) {
      pst = conn.prepareStatement("insert into solar_system_kill_statistical (dt, solar_system_id, kill_quantity, kill_value, create_time) values (?, ?, ?, ?, now()) on duplicate key update kill_quantity = ?, kill_value = ?")
    } else if (classType.getName.equals(classOf[CharacterSink].getName)) {
      pst = conn.prepareStatement("insert into character_kill_statistical (dt, character_id, final_shoot, kill_quantity, kill_value, labels, ship_types, solar_systems, create_time) values (?, ?, ?, ?, ?, ?, ?, ?, now()) on duplicate key update final_shoot = ?, kill_quantity = ?, kill_value = ?, labels = ?, ship_types = ?, solar_systems = ?")
    }
  }

  override def close(): Unit = {
    pst.close()
    conn.close()
  }
}
