package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/global"
	"go_chat/model"
	"go_chat/model/response"
	"strconv"
)

//Login 账号登录检测
func Login(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	accountInt, _ := strconv.Atoi(account)

	var user model.User
	db := global.DB
	db.Where("account = ?", accountInt).Where("password = ?", password).Find(&user)
	if user.Id > 0 {
		//var userRes response.UserResponse
		//userRes.Account = user.Account
		response.OkWithData(user, c)
	} else {
		//response.OkWithData(user, c)
		response.FailWithCodeMessage(response.FORBIDDEN, "账号或密码错误", c)
	}
}

//Enroll 注册账号
func Enroll(c *gin.Context) {
	account := c.PostForm("account")
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	accountInt, _ := strconv.Atoi(account)

	if DuplicateQueryAccount(accountInt) {
		response.FailWithMessage("账号已存在",c)
		return
	}

	db := global.DB

	user := model.User{
		Account:  accountInt,
		Password: password,
		Email:    email,
		Name:     name,
		Phone:    phone,
	}

	db.Create(&user)

	response.OkWithMessage("账号创建成功", c)
}

//Logout 登出账号
func Logout(c *gin.Context) {
	response.Ok(c)
}

//CancelAccount 注销账号
func CancelAccount(c *gin.Context) {
	account := c.PostForm("account")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	var userModel model.User
	db.Where("account = ?", accountInt).Find(&userModel)
	db.Delete(&userModel)
	response.OkWithData("账号注销成功", c)
}

// UpdateUserEmail 修改邮箱
func UpdateUserEmail(c *gin.Context) {
	account := c.PostForm("account")
	Email := c.PostForm("email")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	db.Model(&model.User{}).Where("account = ?", accountInt).Update("email", Email)
	response.OkWithMessage("修改成功", c)
}

// UpdateUserPhone 修改电话
func UpdateUserPhone(c *gin.Context) {
	account := c.PostForm("account")
	Phone := c.PostForm("phone")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	db.Model(&model.User{}).Where("account = ?", accountInt).Update("phone", Phone)
	response.OkWithMessage("修改成功", c)
}

// UpdateUserSignInfo 修改签名
func UpdateUserSignInfo(c *gin.Context) {
	account := c.PostForm("account")
	SignInfo := c.PostForm("sign_info")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	db.Model(&model.User{}).Where("account = ?", accountInt).Update("sign_info", SignInfo)
	response.OkWithMessage("修改成功", c)
}

// UpdateUserPassage 修改密码
func UpdateUserPassage(c *gin.Context) {
	account := c.PostForm("account")
	Passage := c.PostForm("passage")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	db.Model(&model.User{}).Where("account = ?", accountInt).Update("passage", Passage)
	response.OkWithMessage("修改成功", c)
}

// DuplicateTruePassage 验证密码
func DuplicateTruePassage(c *gin.Context) {
	account := c.Query("account")
	Passage := c.Query("passage")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	var user model.User
	db.Where("account = ?", accountInt).Where("passage", Passage).Find(&user)
	if user.Id > 0 {
		response.OkWithMessage("验证通过", c)
	} else {
		response.FailWithMessage("密码不正确", c)
	}
}

// UpdateUserSex 修改性别
func UpdateUserSex(c *gin.Context) {
	account := c.PostForm("account")
	sex := c.PostForm("sex")
	accountInt, _ := strconv.Atoi(account)
	sexInt, _ := strconv.Atoi(sex)
	db := global.DB
	db.Model(&model.User{}).Where("account = ?", accountInt).Update("sex", sexInt)
	response.OkWithMessage("修改成功", c)
}

// UpdateUserName 修改昵称
func UpdateUserName(c *gin.Context) {
	account := c.PostForm("account")
	name := c.PostForm("name")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	db.Model(&model.User{}).Where("account = ?", accountInt).Update("name", name)
	response.OkWithMessage("修改成功", c)
}

// FindUserInformation 返回用户信息
func FindUserInformation(c *gin.Context) {
	account := c.Query("account")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	var userModel model.User
	db.Where("account = ?", accountInt).Find(&userModel)

	userResponse := response.UserResponse{
		Account:  userModel.Account,
		Sex:      userModel.Sex,
		Name:     userModel.Name,
		Phone:    userModel.Phone,
		Email:    userModel.Email,
		SignInfo: userModel.SignInfo,
	}
	response.OkWithData(userResponse, c)
}

// DuplicateQueryAccount 检测账号是否重复
func DuplicateQueryAccount(account int) bool {
	var user model.User
	db := global.DB
	db = db.Where("account = ?", account)
	db.Find(&user)
	if user.Id > 0 {
		return true
	} else {
		return false
	}
}

//s := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
//请求 POST GET
//POST 1.Header c.GetHeader(加密数据)
//     2.Form   c.PostForm("字段名")!!
//GET  1.Query  c.Query("字段名")!!
