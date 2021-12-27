package storage

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"

	"aeon/global"
	"aeon/main/esa-source-zkill/model"
	"aeon/utils"
)

const (
	nilFlag = "invalid"
)

type Kafka struct {
	producer sarama.AsyncProducer
}

func (k *Kafka) Save(msg []byte) {
	global.ESALog.Info("kafka producer receive msg", zap.String("msg", string(msg)))
	kafkaMsgStr := convertKafkaMsgStr(msg)
	global.ESALog.Info("kafka producer send msg", zap.String("msg", kafkaMsgStr))

	kafkaMsg := &sarama.ProducerMessage{
		Topic:     global.ESAConfig.KafkaIn.Topic,
		Value:     sarama.StringEncoder(kafkaMsgStr),
		Timestamp: time.Now(),
	}

	for {
		// TODO 观察是否需要序列化之后再发kafka
		global.ESAKafka.Producer.Input() <- kafkaMsg
		global.ESALog.Info("start monitor kafka producer status...")
		select {
		case success, ok := <-global.ESAKafka.Producer.Successes():
			if ok {
				global.ESALog.Info("Kafka producer success", zap.Any("msg", success))
				return
			}
		case errors, ok := <-global.ESAKafka.Producer.Errors():
			if ok {
				global.ESALog.Error("Kafka producer err", zap.Any("err", errors))
			}
		}
	}
}

func convertKafkaMsgStr(msg []byte) (kafkaMsgStr string) {
	zkill := new(model.ZKill)
	if err := json.Unmarshal(msg, zkill); err != nil {
		global.ESALog.Error("json unmarshal err", zap.Any("err", err))
	}

	/*
		在这里有一些数据是可以直接发送到kafka里的（比如km时间、船型等）。
		也可不发送，由后续数据使用放通过组合killID和hash的方式向ESI发送HTTP请求获取。
		但这样会给下游处理过程造成比较大的影响，处理每一条数据都要发送HTTP请求。
		同时不免偶尔出现网络波动或者ESI服务器无法访问等问题，为了避免，
		在数据源头，就将这些数据做一些简单处理（在不破坏原有数据整体性的情况下）。
	*/

	// timeStamp killId hash solo npc solarSystem iskValue victim attackers labels...
	// iskValue: totalValue#fittedValue#destroyedValue#droppedValue
	// victim: character
	// attackers: character1#character2#character3...
	// character: characterId/AllianceId/CorporationId/Damage/ShipTypeId/FinalBlow

	iskValue := utils.StringSliceBuilder([]string{
		convertToStr(zkill.Zkb.TotalValue),
		"#",
		convertToStr(zkill.Zkb.FittedValue),
		"#",
		convertToStr(zkill.Zkb.DestroyedValue),
		"#",
		convertToStr(zkill.Zkb.DroppedValue),
	})

	victim := convertCharacter(zkill.Killmail.Victim)

	// attackers: character1#character2#character3...
	attackers := ""
	for _, attacker := range zkill.Killmail.Attackers {
		attackers = utils.StringSliceBuilder([]string{attackers, "#", convertCharacter(attacker)})
	}
	attackers = strings.TrimLeft(attackers, "#")

	// timeStamp killId hash solo npc solarSystem iskValue victim attackers labels...
	kafkaMsgStr = utils.StringSliceBuilder([]string{
		convertToStr(zkill.Killmail.KillmailTime.UnixMilli()),
		" ",
		convertToStr(zkill.KillID),
		" ",
		zkill.Zkb.Hash,
		" ",
		convertToStr(zkill.Zkb.Solo),
		" ",
		convertToStr(zkill.Zkb.Npc),
		" ",
		convertToStr(zkill.Killmail.SolarSystemId),
		" ",
		iskValue,
		" ",
		victim,
		" ",
		attackers,
	})

	for _, label := range zkill.Zkb.Labels {
		kafkaMsgStr = utils.StringSliceBuilder([]string{kafkaMsgStr, " ", label})
	}

	return
}

// character: characterId/AllianceId/CorporationId/Damage/ShipTypeId/FinalBlow
func convertCharacter(character model.Character) (str string) {
	// 受害者DamageTaken值大，攻击者DamageDone大
	// 生成的Damage字段选择大的那一个值就行了
	damage := character.DamageDone
	if damage < character.DamageTaken {
		damage = character.DamageTaken
	}
	return utils.StringSliceBuilder([]string{
		convertToStr(character.CharacterId),
		"/",
		convertToStr(character.AllianceId),
		"/",
		convertToStr(character.CorporationId),
		"/",
		convertToStr(damage), // TODO
		"/",
		convertToStr(character.ShipTypeId),
		"/",
		convertToStr(character.FinalBlow),
	})
}

func convertToStr(value interface{}) (converted string) {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.String:
		global.ESALog.Warn("the string does not need to be converted, need update invoking")
		return value.(string)
	case reflect.Int:
		return strconv.Itoa(value.(int))
	case reflect.Int64:
		return strconv.FormatInt(value.(int64), 10)
	case reflect.Float64:
		return strconv.FormatFloat(value.(float64), 'f', 2, 64)
	case reflect.Bool:
		return strconv.FormatBool(value.(bool))
	default:
		global.ESALog.Error("unknown type convert to string", zap.Any("type", valueType.Kind()))
		return nilFlag
	}
}
