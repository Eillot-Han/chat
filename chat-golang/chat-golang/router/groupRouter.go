package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func RegisterGroup(r *gin.RouterGroup) {
	r.GET("/allGroupName", service.AllGroupName)
	r.GET("/userCreateGroup", service.UserCreateGroup)
	r.GET("/userPartGroup",service.UserPartGroup)
	r.GET("/groupLastChatted",service.GroupLastChatted)
	r.POST("/deleteGroup",service.DeleteGroup)
	r.POST("/createGroup",service.CreateGroup)
	r.POST("/updateName",service.UpdateGroupName)
}
