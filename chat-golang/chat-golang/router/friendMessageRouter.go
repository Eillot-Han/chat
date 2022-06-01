package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func RegisterFriendMessage(r *gin.RouterGroup){
	r.GET("/allMessage",service.FindAllFriendMessage)
	r.POST("/addMessage",service.AddFriendMessage)
	r.GET("/inquireMessage",service.InquireFriendMessage)
	r.POST("/deleteMessage",service.DeleteFriendMessage)
	r.GET("/downloadFile",service.DownloadFile)
}
