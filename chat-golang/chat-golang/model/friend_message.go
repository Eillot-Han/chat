package model

type FriendMessage struct {
	Id       int    `gorm:"id" json:"id"`
	RelateId int    `gorm:"relate_id" json:"relate_id"` // 用户的关系id
	FromId   int    `gorm:"from_id" json:"from_id"`     // 发送用户的id
	ToId     int    `gorm:"to_id" json:"to_id"`         // 接收用户的id
	Content  string `gorm:"content" json:"content"`     // 消息内容
	Type     int    `gorm:"type" json:"type"`           // 消息类型 0文件 1文本
	SendTime string `gorm:"send_time" json:"send_time"` // 消息发送时间
	Status   int    `gorm:"status" json:"status"`       // 0正常 1被删除
}

func (FriendMessage) TableName() string {
	return "friend_message"
}
