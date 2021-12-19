package core

import (
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"esa-go-service/global"
)

// Path必须以"/"斜杠结尾
var u = url.URL{Scheme: "wss", Host: "zkillboard.com", Path: "/websocket/"}

type webSocketClient struct {
	clientName string
	Write      chan string
}

func NewWebSocketClient(clientName string) (c *webSocketClient) {
	c = new(webSocketClient)
	c.clientName = clientName
	c.Write = make(chan string)
	return
}

func (c *webSocketClient) Listening(outPut func(msg []byte)) {
	// 建立连接
	// TODO 完善header
	header := map[string][]string{
		"user":    {"1", "2"},
		"email":   {"1"},
		"project": {""},
	}
	client, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Close()
	}()

	// 发送消息
	go func() {
		for writeMsg := range c.Write {
			global.ESALog.Info("WebSocket write message:", zap.String("message", writeMsg))
			if err := client.WriteMessage(websocket.TextMessage, []byte(writeMsg)); err != nil {
				global.ESALog.Error("WebSocket write message err:", zap.Any("err", err))
				// TODO 消息发送失败，可能是网络波动或是其他情况，需要进行处理
				time.Sleep(time.Second * 5)
				c.Write <- `{"action":"sub","channel":"killstream"}`
				continue
			}
		}
	}()

	// TODO
	c.Write <- `{"action":"sub","channel":"killstream"}`

	// 监听消息
	for {
		mt, message, err := client.ReadMessage()
		if err != nil {
			// 消息接收失败，可能是网络波动或是其他情况，需要进行处理
			global.ESALog.Error("WebSocket read message err:", zap.Any("err", err))
			panic(err)
		}
		global.ESALog.Info("WebSocket read message:", zap.Int("mt", mt), zap.ByteString("message", message))
		outPut(message)
	}
}
