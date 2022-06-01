package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_chat/global"
	"go_chat/model"
	"go_chat/model/response"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

// FindAllGroupMessage 返回所有群消息
func FindAllGroupMessage(c *gin.Context) {
	groupId := c.Query("group_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	db := global.DB

	var groupMessageModel []model.GroupMessage
	db.Where("group_id = ?", groupIdInt).Find(&groupMessageModel)

	var groupMessageResponse []response.GroupMessageResponse
	for _, groupMessage := range groupMessageModel {
		group := response.GroupMessageResponse{
			UserId:   groupMessage.UserId,
			Content:  groupMessage.Content,
			Type:     groupMessage.Type,
			SendTime: groupMessage.SendTime,
		}
		groupMessageResponse = append(groupMessageResponse, group)
	}
	response.OkWithData(groupMessageResponse, c)
}

// AddGroupMessage 添加聊天消息
func AddGroupMessage(c *gin.Context) {
	groupId := c.PostForm("group_id")
	userId := c.PostForm("account")
	content := c.PostForm("content")
	Type := c.PostForm("type")
	groupIdInt, _ := strconv.Atoi(groupId)
	userIdInt, _ := strconv.Atoi(userId)
	typeInt, _ := strconv.Atoi(Type)
	t := time.Now().Format("2006-01-02 15:04:05")
	db := global.DB

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
		err = c.SaveUploadedFile(file,"./chat-golang/group/"+groupId+"/file/"+fileNewName)
		if err != nil {
			response.FailWithMessage("文件上传失败", c)
			return
		}
		//将路径保存在content下
		content = "./chat-golang/group/" + groupId + "/file/" + fileNewName
	}

	groupMessage := model.GroupMessage{
		GroupId:  groupIdInt,
		UserId:   userIdInt,
		Content:  content,
		Type:     typeInt,
		SendTime: t,
		Status:   0,
	}

	db.Create(&groupMessage)

	groupMessageResponse := response.GroupMessageResponse{
		UserId:   userIdInt,
		Content:  content,
		Type:     typeInt,
		SendTime: t,
	}

	//更新群最后聊天时间
	db.Model(&model.Group{}).Where("account = ?", groupIdInt).Update("last_chatted", t)
	//更新user最后发言时间
	db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Where("member_id = ?", userIdInt).Update("member_last_chatted", t)
	//更新user发言数量
	db.Model(&model.GroupMember{}).Where("group_id = ?", groupIdInt).Where("member_id = ?", userIdInt).Update("member_send", gorm.Expr("member_send + 1"))
	response.OkWithData(groupMessageResponse, c)
}

// DeleteGroupMessage 删除聊天消息
func DeleteGroupMessage(c *gin.Context) {
	userId := c.PostForm("account")
	groupId := c.PostForm("group_id")
	content := c.PostForm("content")
	sendTime := c.PostForm("send_time")
	userIdInt, _ := strconv.Atoi(userId)
	groupIdInt, _ := strconv.Atoi(groupId)
	db := global.DB
	db.Model(&model.GroupMessage{}).Where("group_id = ?", groupIdInt).Where("content = ?", content).Where("user_id = ?", userIdInt).Where("send_time = ?", sendTime).Update("status", 1)
	response.Ok(c)
}

// InquireGroupMessage 查询聊天信息
func InquireGroupMessage(c *gin.Context) {
	content := c.Query("content")
	groupId := c.Query("group_id")
	groupIdInt, _ := strconv.Atoi(groupId)
	db := global.DB

	var groupMessageModel []model.GroupMessage
	db.Where("group_id = ?", groupIdInt).Where("content LIKE ?", "%"+content+"%").Find(&groupMessageModel)

	if groupMessageModel != nil {
		var groupMessageResponse []response.GroupMessageResponse
		for _, groupMessage := range groupMessageModel {
			if groupMessage.Status == 0 {
				groupMessageRes := response.GroupMessageResponse{
					UserId:   groupMessage.UserId,
					Content:  groupMessage.Content,
					Type:     groupMessage.Type,
					SendTime: groupMessage.SendTime,
				}
				groupMessageResponse = append(groupMessageResponse, groupMessageRes)
			}
		}
		if groupMessageResponse != nil {
			response.OkWithData(groupMessageResponse, c)
		} else {
			response.FailWithCodeMessage(response.FORBIDDEN, "不存在相匹配聊天记录", c)
		}
	} else {
		response.FailWithCodeMessage(response.FORBIDDEN, "不存在相匹配聊天记录", c)
	}
}

// DownloadFile 下载文件
func DownloadFile(c *gin.Context) {
	//获取文件路径
	filePath := c.Query("content")
	filePath, _ = filepath.Abs(filePath)
	fmt.Println(filePath)
	//获取文件内容
	file, _ := os.Open(filePath)
	//获取文件名
	_, fileName := filepath.Split(filePath)
	defer file.Close()

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	//浏览器下载或预览
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	//写出文件
	_, err := io.Copy(c.Writer, file)

	if err != nil {
		response.FailWithCodeMessage(http.StatusInternalServerError, "文件加载失败:"+err.Error(), c)
		return
	}
	response.OkWithData("下载成功", c)
}
