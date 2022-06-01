package main

import (
	"fmt"
	"go_chat/global"
	"go_chat/initialize"
	"os"
	"github.com/go-redis/redis"
)

func main() {
	//连接数据库
	global.DB = initialize.Gorm()
	db, _ := global.DB.DB()
	defer db.Close()

	//连接redis
	Client := redis.NewClient(&redis.Options{
		Addr: 	    "localhost:6379",
		Password: 	"",
		DB: 		0,	
	})
	pong, err := Client.Ping().Result()
	if err != nil {
		fmt.Println("redis connect ping failed")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("redis connect ping")
		fmt.Println(pong)
		global.REDIS = Client
	}

	//项目主程序
	addr := ":8888"
	Router := initialize.Routers()
	Router.Run(addr)
}