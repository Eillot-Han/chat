package model

type User struct {
	Id       int    `gorm:"id" json:"id"`
	Account  int    `gorm:"account" json:"account"`      //账号
	Sex      int    `gorm:"sex" json:"sex"`             // 1男2女0未知
	Name     string `gorm:"name" json:"name"`           // 用户名
	Password string `gorm:"password" json:"password"`   // 密码
	Phone    string `gorm:"phone" json:"phone"`         // 手机号码
	Email    string `gorm:"email" json:"email"`         // email
	SignInfo string `gorm:"sign_info" json:"sign_info"` // 个性签名
}

func (User) TableName() string {
	return "user"
}
