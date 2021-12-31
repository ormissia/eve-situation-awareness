package com.ormissia.zkill.transformation

import com.ormissia.zkill.utils.{Character, ZKillInfo}
import org.apache.flink.api.common.functions.MapFunction

import scala.collection.mutable.ListBuffer

class KafkaLineToZKillInfo extends MapFunction[String, ZKillInfo] {
  override def map(value: String): ZKillInfo = {
    // TODO 异常处理
    val values = value.split(" ")

    val iskValues = values(6).split("#")

    val victimAttributes = values(7).split("/")
    val victim = Character(
      characterId = victimAttributes(0).toInt,
      allianceId = victimAttributes(1).toInt,
      corporationId = victimAttributes(2).toInt,
      damageDone = 0,
      damageTaken = victimAttributes(3).toLong,
      shipTypeId = victimAttributes(4).toInt,
      weaponTypeId = 0,
      finalBlow = false,
    )

    val attackersStr = values(8).split("#")
    val attackers = new ListBuffer[Character]
    attackersStr.foreach(characterStr => {
      //characterStr example: 1406345246/99005338/98496411/7152/639/false
      val AttackerAttributes = characterStr.split("/")
      val attacker = Character(
        characterId = AttackerAttributes(0).toInt,
        allianceId = AttackerAttributes(1).toInt,
        corporationId = AttackerAttributes(2).toInt,
        damageDone = 0,
        damageTaken = AttackerAttributes(3).toLong,
        shipTypeId = AttackerAttributes(4).toInt,
        weaponTypeId = 0,
        finalBlow = AttackerAttributes(5).toBoolean,
      )
      attackers += attacker
    })


    ZKillInfo(
      timeStamp = values(0).toLong,
      killID = values(1).toInt,
      hash = values(2),
      isSolo = values(3).toBoolean,
      killedByNpc = values(4).toBoolean,
      solarSystem = values(5).toInt,
      totalValue = iskValues(0).toDouble,
      fittedValue = iskValues(1).toDouble,
      destroyedValue = iskValues(2).toDouble,
      droppedValue = iskValues(3).toDouble,
      victim = victim,
      attackers = attackers,
      values.drop(9)) // labels为values中去掉前8个固定位置剩下的元素
  }
}
