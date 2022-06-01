package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/global"
	"go_chat/model"
	"go_chat/model/response"
	"gorm.io/gorm"
	"strconv"
)

// GroupAllMember 返回该群所有成员
func GroupAllMember(c *gin.Context) {
	groupId := c.Query("group_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	db := global.DB
	var group model.Group
	//获取群名称
	db.Where("account = ?", groupIdInt).Find(&group)
	groupName := group.Name
	//获取用户id
	var groupMemberModelList []model.GroupMember
	db.Where("group_id = ?", groupIdInt).Find(&groupMemberModelList)
	//获取用户昵称
	var userModelList []model.User
	for _, groupMemberModel := range groupMemberModelList {
		var userModel model.User
		db.Where("account = ?", groupMemberModel.MemberId).Find(&userModel)
		userModelList = append(userModelList, userModel)
	}

	//赋值response
	var groupMemberResponseList []response.GroupMemberResponse
	for _, userModel := range userModelList {
		userId := userModel.Account
		userName := userModel.Name

		groupMemberResponse := response.GroupMemberResponse{
			GroupId:    groupIdInt,
			GroupName:  groupName,
			MemberId:   userId,
			MemberName: userName,
		}
		groupMemberResponseList = append(groupMemberResponseList, groupMemberResponse)
	}
	response.OkWithData(groupMemberResponseList, c)
}

// DeleteGroupMember 删除群成员
func DeleteGroupMember(c *gin.Context) {
	creator := c.PostForm("account")
	groupId := c.PostForm("group_id")
	memberId := c.PostForm("member_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	creatorInt, _ := strconv.Atoi(creator)
	memberIdInt, _ := strconv.Atoi(memberId)
	db := global.DB
	//查询是否为该群管理员
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", creatorInt).Find(&groupMemberModel)
	//存在权限即删除该指定成员用户 总人数减一
	if groupMemberModel.Permissions >= 1 {
		var memberGroup model.GroupMember
		db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Where("member_id = ?", memberIdInt).Find(&memberGroup)
		if memberGroup.Permissions < groupMemberModel.Permissions {
			memberGroup.Status = 1
			memberGroup.Permissions = 0
			db.Save(&memberGroup)
			db.Model(&model.Group{}).Where("account = ?", groupIdInt).Update("user_cnt", gorm.Expr("user_cnt - 1"))
			response.Ok(c)
		} else {
			response.FailWithCodeMessage(response.FORBIDDEN, "您没有权限", c)
		}
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "您没有权限", c)
	}
}

// AddGroupMember 添加群成员
func AddGroupMember(c *gin.Context) {
	creator := c.PostForm("account")
	groupId := c.PostForm("group_id")
	memberId := c.PostForm("member_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	creatorInt, _ := strconv.Atoi(creator)
	memberIdInt, _ := strconv.Atoi(memberId)
	db := global.DB
	//查询是否为该群管理员
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", creatorInt).Find(&groupMemberModel)
	//存在权限即添加该成员
	if groupMemberModel.Permissions >= 1 {
		var groupMember model.GroupMember
		db.Where("group_id = ?", groupIdInt).Where("member_id = ?", memberId).Find(&groupMember)
		if groupMember.Id > 0 && groupMember.Status == 0 {
			response.FailWithMessage("用户已存在", c)
			return
		} else if groupMember.Id > 0 && groupMember.Status == 1 {
			db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Where("member_id = ?", memberIdInt).Update("status", 0)
			response.OkWithData("插入成功", c)
		} else {
			member := model.GroupMember{
				GroupId:  groupIdInt,
				MemberId: memberIdInt,
			}
			db.Create(&member)
			db.Model(&model.Group{}).Where("account = ?", groupIdInt).Update("user_cnt", gorm.Expr("user_cnt + 1"))
			response.OkWithData("插入成功", c)
		}
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "您没有权限", c)
	}
}

// QuitGroup 退出群
func QuitGroup(c *gin.Context) {
	memberId := c.PostForm("account")
	groupId := c.PostForm("group_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	memberIdInt, _ := strconv.Atoi(memberId)
	db := global.DB
	//查询是否为该群群主
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", memberIdInt).Find(&groupMemberModel)
	// 总人数减一
	if groupMemberModel.Permissions < 2 {
		groupMemberModel.Status = 1
		groupMemberModel.Permissions = 0
		db.Save(&groupMemberModel)
		db.Model(&model.Group{}).Where("account = ?", groupIdInt).Update("user_cnt", gorm.Expr("user_cnt - 1"))
		response.Ok(c)
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "请先转移群主", c)
	}
}

// TransferGroupOwnerPermissions 转移群主权限
func TransferGroupOwnerPermissions(c *gin.Context) {
	memberId := c.PostForm("account")
	groupId := c.PostForm("group_id")
	specId := c.PostForm("member_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	specIdInt, _ := strconv.Atoi(specId)
	memberIdInt, _ := strconv.Atoi(memberId)
	db := global.DB
	//查询是否为该群群主
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", memberIdInt).Find(&groupMemberModel)

	if groupMemberModel.Permissions == 2 {
		groupMemberModel.Permissions = 0
		db.Save(&groupMemberModel)
		db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Where("member_id = ?", specIdInt).Update("permissions", 2)
		response.Ok(c)
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "您没有权限", c)
	}
}

// AddGroupAdmin 添加管理员
func AddGroupAdmin(c *gin.Context) {
	memberId := c.PostForm("account")
	groupId := c.PostForm("group_id")
	specId := c.PostForm("member_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	specIdInt, _ := strconv.Atoi(specId)
	memberIdInt, _ := strconv.Atoi(memberId)
	db := global.DB
	//查询是否为该群群主
	var groupMemberModel model.GroupMember
	db.Where("group_id = ?", groupIdInt).Where("member_id = ?", memberIdInt).Find(&groupMemberModel)

	if groupMemberModel.Permissions == 2 {
		db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Where("member_id = ?", specIdInt).Update("permissions", 1)
		response.OkWithMessage("管理员添加成功", c)
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "您没有权限", c)
	}
}
