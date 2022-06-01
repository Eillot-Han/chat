package model

type RelationShip struct {
	Id      int `gorm:"id" json:"id"`
	ShipId  int `gorm:"ship_id" json:"ship_id"`   // 用户关系id
	SmallId int `gorm:"small_id" json:"small_id"` // 用户A的id
	BigId   int `gorm:"big_id" json:"big_id"`     // 用户B的id
	Status  int `gorm:"status" json:"status"`   // 用户:0-正常, 1-删除
}

func (RelationShip) TableName() string {
	return "relation_ship"
}
