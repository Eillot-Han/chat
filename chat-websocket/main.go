package main

import (
	"chat-websocket/websocket"
	"fmt"
	"net/http"

)

//进程
var process = websocket.NewProcess()

func serveWs(process *websocket.Process, w http.ResponseWriter, r *http.Request) {
    //从请求中解析出聊天室和用户信息
    r.ParseForm()
    roomid := r.Form["roomid"][0]
    uid := r.Form["uid"][0]

    fmt.Println("WebSocket Endpoint Hit, room = ", roomid, ",uid=", uid)
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    if process.Pools[roomid] == nil {
        pool := websocket.NewPool()
        go pool.Start()
        process.Pools[roomid] = pool
    }

    client := &websocket.Client{
        Conn: conn,
        Pool: process.Pools[roomid],
    }

    process.Pools[roomid].Register <- client
    client.Read()
}

func setupRoutes() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        serveWs(process, w, r)
    })
}

func main() {
    fmt.Println("Distributed Chat App")
    setupRoutes()
    //监听8080端口
    http.ListenAndServe(":8080", nil)
}
