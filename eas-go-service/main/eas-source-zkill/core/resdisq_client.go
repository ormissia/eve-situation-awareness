package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"

	"eas-go-service/global"
)

type redisQClient struct {
	clientName string
}

func NewRedisQClient(clientName string) (c *redisQClient) {
	c = new(redisQClient)
	c.clientName = clientName
	return
}

func (c *redisQClient) Listening(outPut func(msg []byte)) {
	defer func() {
		if err := recover(); err != nil {
			global.EASLog.Error("Listening err: ", zap.Any("err", err))
		}
	}()

	// 定时器
	tick := time.Tick(time.Second)
	url := fmt.Sprintf("https://redisq.zkillboard.com/listen.php?queueID=%s&ttw=3", c.clientName)

	client := GetRedisQClient(url)

	for {
		<-tick

		msg, err := client()
		if err != nil {
			global.EASLog.Error("RedisQ client get msg err", zap.Any("err", err))
			continue
		}

		res := make(map[string]interface{})
		if err = json.Unmarshal(msg, &res); err != nil {
			global.EASLog.Error("Json unmarshal err", zap.Any("err", err))
			continue
		}

		// 调用回调函数
		packageMap, ok := res["package"]
		if ok && packageMap != nil {
			bytes, _ := json.Marshal(packageMap)
			outPut(bytes)
		} else {
			global.EASLog.Info("RedisQ no data")
		}
	}
}

func GetRedisQClient(url string) func() (msg []byte, err error) {
	return func() (msg []byte, err error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = resp.Body.Close()
		}()
		return ioutil.ReadAll(resp.Body)
	}
}
