package com.ormissia.zkill.utils

import scala.collection.mutable.ListBuffer

// 0         1      2    3    4   5           6        7      8         9
// timeStamp killId hash solo npc solarSystem iskValue victim attackers labels...
//           0          1           2              3
// iskValue: totalValue#fittedValue#destroyedValue#droppedValue
// victim: character
// attackers: character1#character2#character3...
// character: characterId/AllianceId/CorporationId/Damage/ShipTypeId/FinalBlow
case class ZKillInfo(
                      timeStamp: Long,
                      killID: Int,
                      hash: String,
                      isSolo: Boolean,
                      killedByNpc: Boolean,
                      solarSystem: Int,
                      totalValue: Double,
                      fittedValue: Double,
                      destroyedValue: Double,
                      droppedValue: Double,
                      victim: Character,
                      attackers: ListBuffer[Character],
                      labels: Array[String],
                    )

/*
// go struct
type Character struct {
    CharacterId    int     `json:"character_id"`
    AllianceId     int     `json:"alliance_id"`
    CorporationId  int     `json:"corporation_id"`
    DamageDone     int     `json:"damage_done"`
    DamageTaken    int     `json:"damage_taken"`
    ShipTypeId     int     `json:"ship_type_id"`
    WeaponTypeId   int     `json:"weapon_type_id"`
    SecurityStatus float64 `json:"security_status"`
    FinalBlow      bool    `json:"final_blow"`
}
*/
// character: characterId/AllianceId/CorporationId/Damage/ShipTypeId/FinalBlow
case class Character(
                      characterId: Int,
                      allianceId: Int,
                      corporationId: Int,
                      damageDone: Long,
                      damageTaken: Long,
                      shipTypeId: Int,
                      weaponTypeId: Int,
                      finalBlow: Boolean,
                    )


case class SolarSystemSink(
                            dt: Int,
                            solarSystemId: Int,
                            killQuantity: Int,
                            killValue: Double,
                          )

object GlobalConstants {

}
