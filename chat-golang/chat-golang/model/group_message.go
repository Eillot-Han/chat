package model

type GroupMessage struct {
	Id       int    `gorm:"id" json:"id"`
	GroupId  int    `gorm:"group_id" json:"group_id"`   // 用户的群id
	UserId   int    `gorm:"user_id" json:"user_id"`     // 发送用户的id
	Content  string `gorm:"content" json:"content"`     // 消息内容
	Type     int    `gorm:"type" json:"type"`           // 消息类型 0文件 1文本
	SendTime string `gorm:"send_time" json:"send_time"` // 消息发送时间
	Status   int    `gorm:"status" json:"status"`       // 0正常 1被删除
}

func (GroupMessage) TableName() string {
	return "group_message"
}
