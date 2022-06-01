package initialize

import (
	"go_chat/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()

	user := Router.Group("user")
	{
		router.RegisterUser(user)
	}
	relationship := Router.Group("relationship")
	{
		router.RegisterRelationship(relationship)
	}
	group := Router.Group("group")
	{
		router.RegisterGroup(group)
	}
	groupMessage := Router.Group("groupMessage")
	{
		router.RegisterGroupMessage(groupMessage)
	}
	groupMember := Router.Group("groupMember")
	{
		router.RegisterGroupMember(groupMember)
	}
	friendMessage := Router.Group("friendMessage")
	{
		router.RegisterFriendMessage(friendMessage)
	}
	return Router
}