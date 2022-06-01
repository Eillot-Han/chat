package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/global"
	"go_chat/model"
	"go_chat/model/response"
	"math/rand"
	"strconv"
	"time"
)

// AllGroupName 返回所有群组的id和名称
func AllGroupName(c *gin.Context) {
	var groupMap map[int]string
	db := global.DB

	var groupModelList []model.Group
	db.Find(&groupModelList)

	groupMap = make(map[int]string)

	for _, groupModel := range groupModelList {
		name := groupModel.Name
		id := groupModel.Id
		groupMap[id] = name
	}

	response.OkWithData(groupMap, c)
}

// UserCreateGroup 返回所拥有的群
func UserCreateGroup(c *gin.Context) {
	creator := c.Query("account")
	creatorInt,_ := strconv.Atoi(creator)

	db := global.DB
	var groupModeList []model.Group
	db.Where("creator = ?", creatorInt).Find(&groupModeList)
	//赋值response
	var groupResponseMap []response.GroupResponse
	for _, groupModel := range groupModeList {
		account := groupModel.Account
		name := groupModel.Name
		userCnt := groupModel.UserCnt

		groupResponse := response.GroupResponse{
			Account: account,
			Name:    name,
			UserCnt: userCnt,
		}

		groupResponseMap = append(groupResponseMap, groupResponse)
	}

	response.OkWithData(groupResponseMap, c)
}

// UserPartGroup 返回所参与的群
func UserPartGroup(c *gin.Context) {
	userId := c.Query("account")
	userIdInt,_ := strconv.Atoi(userId)
	db := global.DB

	var groupMemberModelList []model.GroupMember
	db.Where("member_id = ?", userIdInt).Find(&groupMemberModelList)

	var groupModelList []model.Group
	for _, groupMemberModel := range groupMemberModelList {
		//判断群组是否已经删除
		if groupMemberModel.Status == 0 {
			var groupModel model.Group
			db.Where("account = ?", groupMemberModel.GroupId).Find(&groupModel)
			groupModelList = append(groupModelList, groupModel)
		}
	}
	//赋值response
	var groupResponseList []response.GroupResponse
	for _, groupModel := range groupModelList {
		account := groupModel.Account
		name := groupModel.Name
		userCnt := groupModel.UserCnt

		groupResponse := response.GroupResponse{
			Account: account,
			Name:    name,
			UserCnt: userCnt,
		}

		groupResponseList = append(groupResponseList, groupResponse)
	}
	response.OkWithData(groupResponseList, c)
}

// GroupLastChatted 返回群组最后聊天时间
func GroupLastChatted(c *gin.Context) {
	groupId := c.Query("group_id")
	groupIdInt,_ := strconv.Atoi(groupId)
	db := global.DB
	var groupModel model.Group
	db.Where("account = ?", groupIdInt).Find(&groupModel)

	groupLastChatted := response.GroupLastChattedResponse{
		Account:     groupModel.Account,
		Name:        groupModel.Name,
		LastChatted: groupModel.LastChatted,
	}

	response.OkWithData(groupLastChatted, c)
}

// DeleteGroup 群主删除群组
func DeleteGroup(c *gin.Context) {
	creator := c.PostForm("account")
	groupId := c.PostForm("group_id")
	creatorInt,_ := strconv.Atoi(creator)
	groupIdInt,_ := strconv.Atoi(groupId)
	db := global.DB
	//查询是否为该群群主
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", creatorInt).Find(&groupMemberModel)

	//将群组以及群成员标注为已删除 标注群人数为0
	if groupMemberModel.Permissions == 2 {
		db.Model(&model.Group{}).Where("account = ?", groupIdInt).Updates(map[string]interface{}{"status": 1, "user_cnt": 0})
		db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Update("status", 1)
	}

	response.OkWithMessage("群组删除成功",c)
}

// CreateGroup 创建群
func CreateGroup(c *gin.Context) {
	creator := c.PostForm("account")
	creatorInt,_ := strconv.Atoi(creator)
	var account int
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		account = r.Intn(1000000000)
		if !DuplicateQueryGroupAccount(account) {
			break
		}
	}

	db := global.DB
	//添加群信息
	group := model.Group{
		Account: account,
		Name:    "新建群聊",
		Creator: creatorInt,
		UserCnt: 1,
		Status:  0,
	}
	db.Create(&group)

	member := model.GroupMember{
		GroupId:  account,
		MemberId: creatorInt,
		Permissions: 2,
		Status:   0,
	}
	db.Create(&member)

	groupResponse := response.GroupResponse{
		Account: account,
		Name: group.Name,
		UserCnt: group.UserCnt,
	}

	response.OkWithData(groupResponse,c)
}

// UpdateGroupName 修改群名称
func UpdateGroupName(c *gin.Context) {
	creator := c.PostForm("account")
	groupId := c.PostForm("group_id")
	name := c.PostForm("name")
	creatorInt,_ := strconv.Atoi(creator)
	groupIdInt,_ := strconv.Atoi(groupId)
	db := global.DB
	//查询是否为该群管理员
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", creatorInt).Find(&groupMemberModel)

	if groupMemberModel.Permissions >= 1 {
		db.Model(&model.Group{}).Where("account = ?", groupIdInt).Update("name",name)
		response.OkWithMessage("群名称修改成功",c)
	}else{
		response.FailWithCodeMessage(response.FORBIDDEN,"您没有权限",c)
	}
}

// DuplicateQueryGroupAccount 检测账号是否重复
func DuplicateQueryGroupAccount(account int) bool {
	var group model.Group
	db := global.DB
	db.Where("account = ?", account).Find(&group)
	if group.Id > 0 {
		return true
	} else {
		return false
	}

}
