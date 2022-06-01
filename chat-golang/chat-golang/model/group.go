package model

type Group struct {
	Id          int    `gorm:"id" json:"id"`
	Account     int    `gorm:"account" json:"account"`           // 群id
	Name        string `gorm:"name" json:"name"`                 // 群名称
	Creator     int    `gorm:"creator" json:"creator"`           // 创建者用户id
	UserCnt     int    `gorm:"user_cnt" json:"user_cnt"`         // 成员人数
	Status      int    `gorm:"status" json:"status"`             // 是否删除,0-正常，1-删除
	LastChatted string `gorm:"last_chatted" json:"last_chatted"` // 最后聊天时间
}

func (Group) TableName() string {
	return "group"
}
