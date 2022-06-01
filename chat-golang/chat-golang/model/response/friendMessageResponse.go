package response

type FriendMessageResponse struct {
	FromId   int    `gorm:"from_id" json:"from_id"`     // 发送用户的id
	ToId     int    `gorm:"to_id" json:"to_id"`         // 接收用户的id
	Content  string `gorm:"content" json:"content"`     // 消息内容
	Type     int    `gorm:"type" json:"type"`           // 消息类型 0文件 1文本
	SendTime string `gorm:"send_time" json:"send_time"` // 消息发送时间
}
