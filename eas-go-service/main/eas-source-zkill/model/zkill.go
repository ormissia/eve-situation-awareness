package model

import "time"

type ZKill struct {
	KillID   int `json:"killID"`
	Killmail struct {
		Attackers     []Character `json:"attacers"`
		KillmailTime  time.Time   `json:"killmail_time"`
		SolarSystemId int         `json:"solar_system_id"`
		Victim        Character   `json:"victim"`
	} `json:"killmail"`
	Zkb struct {
		Awox           bool     `json:"awox"`
		DestroyedValue float64  `json:"destroyedValue"`
		DroppedValue   float64  `json:"droppedValue"`
		FittedValue    float64  `json:"fittedValue"`
		Hash           string   `json:"hash"`
		Href           string   `json:"href"`
		Labels         []string `json:"labels"`
		LocationID     int      `json:"locationID"`
		Npc            bool     `json:"npc"`
		Points         int      `json:"points"`
		Solo           bool     `json:"solo"`
		TotalValue     float64  `json:"totalValue"`
	} `json:"zkb"`
}

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
