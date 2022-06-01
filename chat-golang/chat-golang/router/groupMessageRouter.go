package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func RegisterGroupMessage(r *gin.RouterGroup){
	r.GET("/allMessage",service.FindAllGroupMessage)
	r.POST("/addMessage",service.AddGroupMessage)
	r.POST("/deleteMessage",service.DeleteGroupMessage)
	r.GET("/inquireMessage",service.InquireGroupMessage)
	r.GET("/download",service.DownloadFile)
}