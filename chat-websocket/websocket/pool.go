package websocket

import (
	"fmt"
)

// Pool 线程池
//包含发送通信所需的channels
type Pool struct {
	Register   chan *Client     //传递客户端连接信息
	Unregister chan *Client     //传递客户端断开连接
	Clients    map[*Client]bool //判断客户端活跃与否
	Broadcast  chan Message
}

// NewPool 新建线程池
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		//随机执行case语句
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			for client, _ := range pool.Clients {
				_ = client.Conn.WriteJSON(Message{Type: 1, Body: "New user join..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			for client, _ := range pool.Clients {
				_ = client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected"})
			}
			break
		case message := <-pool.Broadcast:
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
			break
		}
	}
}
