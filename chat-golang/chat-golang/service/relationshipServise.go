package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/global"
	"go_chat/model"
	"go_chat/model/response"
	"math/rand"
	"strconv"
	"time"
)

// RelationshipAll 返回用户关系
func RelationshipAll(c *gin.Context) {
	account := c.Query("account")
	accountInt, _ := strconv.Atoi(account)
	db := global.DB
	var relationshipModelList []model.RelationShip
	db.Where("small_id = ?", accountInt).Find(&relationshipModelList)
	//赋值response
	var relationshipResponseList []response.RelationShipResponse
	for _, relationshipModel := range relationshipModelList {
		shipId := relationshipModel.ShipId
		smellId := relationshipModel.SmallId
		bigId := relationshipModel.BigId
		status := relationshipModel.Status
		relationResponse := response.RelationShipResponse{
			ShipId: shipId,
			SmallId: smellId,
			BigId:   bigId,
			Status:  status,
		}
		relationshipResponseList = append(relationshipResponseList, relationResponse)
	}
	response.OkWithData(relationshipResponseList, c)
}

// DeleteRelationship 删除好友
func DeleteRelationship(c *gin.Context) {
	account := c.PostForm("account")
	friendId := c.PostForm("friend_id")
	accountInt, _ := strconv.Atoi(account)
	friendIdInt, _ := strconv.Atoi(friendId)
	db := global.DB
	db.Model(&model.RelationShip{}).Where("small_id = ?", accountInt).Where("big_id = ?", friendIdInt).Update("status", 1)
	db.Model(&model.RelationShip{}).Where("small_id = ?", friendIdInt).Where("big_id = ?", accountInt).Update("status", 1)
	response.OkWithMessage("好友删除成功", c)
}

// AddRelationship 添加好友
func AddRelationship(c *gin.Context) {
	account := c.PostForm("account")
	friendId := c.PostForm("friend_id")
	accountInt, _ := strconv.Atoi(account)
	friendIdInt, _ := strconv.Atoi(friendId)
	db := global.DB
	var relationshipModel model.RelationShip
	db.Where("small_id = ?", accountInt).Where("big_id = ?", friendIdInt).Find(&relationshipModel)
	//判断曾经是否是好友
	if relationshipModel.Id > 0 && relationshipModel.Status == 1 {
		db.Model(&model.RelationShip{}).Where("small_id = ?", accountInt).Where("big_id = ?", friendIdInt).Update("status", 0)
		db.Model(&model.RelationShip{}).Where("small_id = ?", friendIdInt).Where("big_id = ?", accountInt).Update("status", 0)
		response.OkWithMessage("好友添加成功", c)
	} else if relationshipModel.Id > 0 && relationshipModel.Status == 0 {
		response.FailWithCodeMessage(response.FORBIDDEN, "已为好友，无需再次添加", c)
		return
	} else {
		//查询shipId是否重复
		var shipId int
		for {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			shipId = r.Intn(1000000000)
			if !DuplicateQueryShipId(shipId) {
				break
			}
		}

		relationship := model.RelationShip{
			ShipId:  shipId,
			SmallId: accountInt,
			BigId:   friendIdInt,
			Status:  0,
		}
		db.Create(&relationship)

		Relationship := model.RelationShip{
			ShipId:  shipId,
			SmallId: friendIdInt,
			BigId:   accountInt,
			Status:  0,
		}
		db.Create(&Relationship)
		if db.Error != nil {
			response.FailWithMessage("好友添加失败", c)
		}
		response.OkWithMessage("好友添加成功", c)
	}
}

// DuplicateQueryShipId 查询好友关系是否重复
func DuplicateQueryShipId(shipId int) bool {
	var relationship model.RelationShip
	db := global.DB
	db = db.Where("shipId = ?", shipId)
	db.Find(&relationship)
	if relationship.Id > 0 {
		return true
	} else {
		return false
	}
}
