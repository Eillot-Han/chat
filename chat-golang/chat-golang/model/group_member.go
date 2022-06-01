package model

type GroupMember struct {
	Id                int    `gorm:"id" json:"id"`
	GroupId           int    `gorm:"group_id" json:"group_id"`                       // 群id
	MemberId          int    `gorm:"member_id" json:"member_id"`                     // 成员id
	MemberUnread      int    `gorm:"member_unread" json:"member_unread"`             // 未读消息数量
	MemberSend        int    `gorm:"member_send" json:"member_send"`                 // 用户发送消息数量
	MemberLastChatted string `gorm:"member_last_chatted" json:"member_last_chatted"` // 用户最后聊天时间
	Status            int    `gorm:"status" json:"status"`                           // 是否删除,0-正常，1-删除
	Permissions       int    `gorm:"permissions" json:"permissions"`                 // 判断权限，0-正常，1-管理员，2-群主
}

func (GroupMember) TableName() string {
	return "group_member"
}
