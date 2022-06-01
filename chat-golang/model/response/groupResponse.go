package response

type GroupResponse struct {
	Account int    `json:"account"`  // 群id
	Name    string `json:"name"`     // 群名称
	UserCnt int    `json:"user_cnt"` // 成员人数
}

type GroupLastChattedResponse struct {
	Account     int    `json:"account"`      // 群id
	Name        string `json:"name"`         // 群名称
	LastChatted string `json:"last_chatted"` // 最后聊天时间
}
