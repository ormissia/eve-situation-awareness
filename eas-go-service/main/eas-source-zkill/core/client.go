package core

import (
	"net/url"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"eas-go-service/global"
)

func NewClient(readBufferSize int) *Client {
	return &Client{
		Read:  make(chan []byte, readBufferSize),
		Write: make(chan string),
	}
}

type Client struct {
	Read  chan []byte
	Write chan string
}

// Path必须以"/"斜杠结尾
var u = url.URL{Scheme: "wss", Host: "zkillboard.com", Path: "/websocket/"}

func (client *Client) Connect() {
	// 建立连接
	// TODO 完善header
	header := map[string][]string{
		"user":    {"1", "2"},
		"email":   {"1"},
		"project": {""},
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = c.Close()
		close(client.Write)
		close(client.Read)
	}()

	// 发送消息
	go func() {
		for writeMsg := range client.Write {
			global.EASLog.Info("WebSocket write message:", zap.String("message", writeMsg))
			if err := c.WriteMessage(websocket.TextMessage, []byte(writeMsg)); err != nil {
				global.EASLog.Error("WebSocket write message err:", zap.Any("err", err))
				// TODO 消息发送失败，可能是网络波动或是其他情况，需要进行处理
				continue
			}
		}
	}()

	// 监听消息
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			// 消息接收失败，可能是网络波动或是其他情况，需要进行处理
			global.EASLog.Error("WebSocket read message err:", zap.Any("err", err))
			panic(err)
		}
		global.EASLog.Info("WebSocket read message:", zap.Int("mt", mt), zap.ByteString("message", message))
		client.Read <- message
	}
}
