package response

type GroupMemberResponse struct {
	GroupId    int    `json:"group_id"`    // 群id
	GroupName  string `json:"group_name"`  // 群名称
	MemberId   int    `json:"member_id"`   // 成员id
	MemberName string `json:"member_name"` // 成员名称
}
