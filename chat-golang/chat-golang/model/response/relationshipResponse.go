package response

type RelationShipResponse struct {
	ShipId  int `json:"ship_id"`  // 好友关系id
	SmallId int `json:"small_id"` // 用户A的id
	BigId   int `json:"big_id"`   // 用户B的id
	Status  int `json:"status"`   // 用户:0-正常, 1-删除
}
