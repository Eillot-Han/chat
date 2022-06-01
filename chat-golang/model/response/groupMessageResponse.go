package response

type GroupMessageResponse struct {
	UserId   int    `gorm:"user_id" json:"user_id"`     // 发送用户的id
	Content  string `gorm:"content" json:"content"`     // 消息内容
	Type     int    `gorm:"type" json:"type"`           // 消息类型 0文件 1文本
	SendTime string `gorm:"send_time" json:"send_time"` // 消息发送时间
}
