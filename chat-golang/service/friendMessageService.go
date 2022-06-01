package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/global"
	"go_chat/model"
	"go_chat/model/response"
	"path"
	"strconv"
	"time"
)

// FindAllFriendMessage 返回与当前好友的所有聊天记录
func FindAllFriendMessage(c *gin.Context) {
	shipId := c.Query("ship_id")
	shipIdInt, _ := strconv.Atoi(shipId)
	db := global.DB

	var friendMessageModel []model.FriendMessage
	db.Where("relate_id = ?", shipIdInt).Find(&friendMessageModel)

	//赋值response
	if friendMessageModel != nil {
		var friendMessageResponse []response.FriendMessageResponse
		for _, friendMessage := range friendMessageModel {
			if friendMessage.Status == 0 {
				friendMessageRes := response.FriendMessageResponse{
					FromId:   friendMessage.FromId,
					ToId:     friendMessage.ToId,
					Content:  friendMessage.Content,
					Type:     friendMessage.Type,
					SendTime: friendMessage.SendTime,
				}
				friendMessageResponse = append(friendMessageResponse, friendMessageRes)
			}
		}
		if friendMessageResponse != nil {
			response.OkWithData(friendMessageResponse, c)
		} else {
			response.FailWithCodeMessage(response.FORBIDDEN, "不存在与该用户的聊天记录", c)
		}
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "不存在与该用户的聊天记录", c)
	}
}

// AddFriendMessage 添加聊天记录
func AddFriendMessage(c *gin.Context) {
	fromID := c.PostForm("account")
	shipId := c.PostForm("ship_id")
	content := c.PostForm("content")
	_type := c.PostForm("type")
	fromIDInt, _ := strconv.Atoi(fromID)
	shipIdInt, _ := strconv.Atoi(shipId)
	typeInt, _ := strconv.Atoi(_type)
	t := time.Now().Format("2006-01-02 15:04:05")
	db := global.DB

	var friendship model.RelationShip
	db.Where("small_id = ?", fromIDInt).Where("ship_id = ?", shipIdInt).Find(&friendship)

	if typeInt == 0 {
		file, err := c.FormFile("upload_file")
		if err != nil {
			response.FailWithMessage("文件获取失败", c)
			return
		}
		//获取文件后缀
		ext := path.Ext(file.Filename)
		//定义文件新名字  当前时间加后缀
		fileNewName := time.Now().Format("20060102150405") + strconv.Itoa(time.Now().Nanosecond()) + ext
		//保存文件到路径下
		err = c.SaveUploadedFile(file,"./chat-golang/friend/"+shipId+"/file/"+fileNewName)
		if err != nil {
			response.FailWithMessage("文件上传失败", c)
			return
		}
		//将路径保存在content下
		content = "./chat-golang/friend/" + shipId + "/file/" + fileNewName
	}

	friendMessageModel := model.FriendMessage{
		RelateId: friendship.ShipId,
		FromId:   friendship.SmallId,
		ToId:     friendship.BigId,
		Content:  content,
		Type:     typeInt,
		SendTime: t,
		Status:   0,
	}

	db.Create(&friendMessageModel)

	friendMessageResponse := response.FriendMessageResponse{
		FromId:   friendship.SmallId,
		ToId:     friendship.BigId,
		Content:  content,
		Type:     typeInt,
		SendTime: t,
	}

	response.OkWithData(friendMessageResponse, c)
}

// InquireFriendMessage 查询聊天记录
func InquireFriendMessage(c *gin.Context) {
	shipId := c.Query("ship_id")
	content := c.Query("content")
	shipIdInt, _ := strconv.Atoi(shipId)
	db := global.DB

	var friendMessageModel []model.FriendMessage
	db.Debug().Where("relate_id = ?", shipIdInt).Where("content LIKE ?", "%"+content+"%").Find(&friendMessageModel)

	if friendMessageModel != nil {
		var friendMessageResponse []response.FriendMessageResponse
		for _, friendMessage := range friendMessageModel {
			if friendMessage.Status == 0 {
				friendMessageRes := response.FriendMessageResponse{
					FromId:   friendMessage.FromId,
					ToId:     friendMessage.ToId,
					Content:  friendMessage.Content,
					Type:     friendMessage.Type,
					SendTime: friendMessage.SendTime,
				}
				friendMessageResponse = append(friendMessageResponse, friendMessageRes)
			}
		}
		if friendMessageResponse != nil {
			response.OkWithData(friendMessageResponse, c)
		} else {
			response.FailWithCodeMessage(response.FORBIDDEN, "不存在相匹配聊天记录", c)
		}
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "不存在相匹配聊天记录", c)
	}
}

// DeleteFriendMessage 删除聊天记录
func DeleteFriendMessage(c *gin.Context) {
	shipId := c.PostForm("ship_id")
	content := c.PostForm("content")
	sendTime := c.PostForm("send_time")
	shipIdInt, _ := strconv.Atoi(shipId)
	db := global.DB

	db.Model(&model.FriendMessage{}).Where("relate_id = ?", shipIdInt).Where("content = ?", content).Where("send_time = ?", sendTime).Update("status", 1)

	response.Ok(c)
}
