package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func RegisterGroupMember(r *gin.RouterGroup) {
	r.GET("/allMember", service.GroupAllMember)
 	r.POST("/deleteMember",service.DeleteGroupMember)
	r.POST("/addMember",service.AddGroupMember)
	r.POST("/quit",service.QuitGroup)
	r.POST("/transferGroupOwnerPermissions",service.TransferGroupOwnerPermissions)
	r.POST("/addAdmin",service.AddGroupAdmin)
}
