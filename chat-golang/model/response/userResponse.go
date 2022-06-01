package response

type UserResponse struct {
	Account  int    `json:"account"`   // 账号
	Sex      int    `json:"sex"`       // 1男2女0未知
	Name     string `json:"name"`      // 用户名
	Phone    string `json:"phone"`     // 手机号码
	Email    string `json:"email"`     // email
	SignInfo string `json:"sign_info"` // 个性签名
}
