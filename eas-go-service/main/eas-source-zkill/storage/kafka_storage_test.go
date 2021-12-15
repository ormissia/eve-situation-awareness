package storage

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestConvertStr(t *testing.T) {
	a := 9687622.36
	fmt.Printf("%T\n", a)
	fmt.Println(convertToStr(a))
	fmt.Println(strconv.FormatFloat(9687622.36, 'f', 2, 64))
	var b int64
	b = 123234234
	fmt.Println(convertToStr(b))
}

const (
	zkillMsgJSON = `
{
    "killID":97381235,
    "killmail":{
        "attackers":[
            {
                "alliance_id":99005338,
                "character_id":2113167159,
                "corporation_id":98110767,
                "damage_done":25553,
                "final_blow":true,
                "security_status":5,
                "ship_type_id":3756,
                "weapon_type_id":3146
            }
        ],
        "killmail_id":97381235,
        "killmail_time":"2021-12-15T15:30:19Z",
        "solar_system_id":30002825,
        "victim":{
            "alliance_id":99003581,
            "character_id":2113984713,
            "corporation_id":98540583,
            "damage_taken":25619,
            "items":[
                {
                    "flag":5,
                    "item_type_id":61130,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":13,
                    "item_type_id":4405,
                    "quantity_dropped":1,
                    "singleton":0
                },
                {
                    "flag":19,
                    "item_type_id":3244,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":92,
                    "item_type_id":31718,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":20,
                    "item_type_id":12056,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":12,
                    "item_type_id":4405,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":27,
                    "item_type_id":13003,
                    "quantity_dropped":1,
                    "singleton":0
                },
                {
                    "flag":5,
                    "item_type_id":61134,
                    "quantity_dropped":1,
                    "singleton":0
                },
                {
                    "flag":11,
                    "item_type_id":2048,
                    "quantity_dropped":1,
                    "singleton":0
                },
                {
                    "flag":93,
                    "item_type_id":31790,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":5,
                    "item_type_id":448,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":87,
                    "item_type_id":2185,
                    "quantity_destroyed":3,
                    "singleton":0
                },
                {
                    "flag":134,
                    "item_type_id":57028,
                    "quantity_dropped":4,
                    "singleton":0
                },
                {
                    "flag":28,
                    "item_type_id":13003,
                    "quantity_destroyed":1,
                    "singleton":0
                },
                {
                    "flag":94,
                    "item_type_id":31790,
                    "quantity_destroyed":1,
                    "singleton":0
                }
            ],
            "position":{
                "x":-3818223370404.8604,
                "y":2185311895247.3003,
                "z":910126645233.214
            },
            "ship_type_id":17480
        }
    },
    "zkb":{
        "awox":false,
        "destroyedValue":48753961.48,
        "droppedValue":3427801.16,
        "fittedValue":47548378.47,
        "hash":"8cb19c0934a4edd520fb0ee6f688bc5c0d16b235",
        "href":"https://esi.evetech.net/v1/killmails/97381235/8cb19c0934a4edd520fb0ee6f688bc5c0d16b235/",
        "labels":[
            "cat:6",
            "solo",
            "pvp",
            "nullsec"
        ],
        "locationID":50014208,
        "npc":false,
        "points":41,
        "solo":true,
        "totalValue":52181762.64
    }
}
`
)

func TestConvertKafkaMsgStr(t *testing.T) {
	kafkaMsgStr := convertKafkaMsgStr([]byte(zkillMsgJSON))
	// killID hash solo npc totalValue fittedValue destroyedValue droppedValue labels...
	msgs := strings.Split(kafkaMsgStr, " ")
	fmt.Println(msgs)
	if len(msgs) < 8 {
		t.Error("missing values")
	}
}
