package websocket

import (
	"github.com/gorilla/websocket"
)

// Client 客户端结构体
type Client struct {
	ID      string
	Conn    *websocket.Conn
	Pool    *Pool
	Process *Process
}

// Message 一条信息
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// Read 监听client的websocket连接上发出的消息
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	//一旦收到消息
	//把消息传递到线程池的broadcast channel
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
	}
}
