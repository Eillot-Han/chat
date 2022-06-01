package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func RegisterUser(r *gin.RouterGroup) {
	r.GET("/login", service.Login)
	r.POST("/enroll", service.Enroll)
	r.GET("/logout", service.Logout)
	r.POST("/cancelAccount",service.CancelAccount)
	r.POST("/updateEmail",service.UpdateUserEmail)
	r.POST("/updatePhone",service.UpdateUserPhone)
	r.POST("/updateSignInfo",service.UpdateUserSignInfo)
	r.POST("/updatePassage",service.UpdateUserPassage)
	r.POST("/updateSex",service.UpdateUserSex)
	r.POST("/updateName",service.UpdateUserName)
	r.GET("/duplicatePassage",service.DuplicateTruePassage)
	r.GET("/findInformation",service.FindUserInformation)
}
