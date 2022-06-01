package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func RegisterRelationship(r *gin.RouterGroup){
	r.GET("/allFriend",service.RelationshipAll)
	r.POST("/addFriend",service.AddRelationship)
	r.POST("/deleteFriend",service.DeleteRelationship)
}