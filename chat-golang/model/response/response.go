package response

import (
	"encoding/json"
	"go_chat/global"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS   = 200
	FORBIDDEN = 403
	ERROR     = 500
)

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	DateTime  string      `json:"date_time"`
	TimeStamp int64       `json:"time_stamp"`
	Sign      string      `json:"sign,omitempty"`
}

type RequestMap struct {
	Uri       string    `json:"uri"`
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	response := Response{
		Code:      code,
		Data:      data,
		Msg:       msg,
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TimeStamp: time.Now().Unix(),
	}
	res := &response
	if code != SUCCESS {
		res.Msg = res.Msg + " (" + strconv.Itoa(res.Code) + ")"
	}
	redis := global.REDIS
	requestMap := RequestMap{
		Uri:       c.Request.URL.String(),
		Ip:        c.GetString("ip"),
		CreatedAt: time.Time{},
	}

	bytes, _ := json.Marshal(requestMap)
	redis.LPush("queue_visit", string(bytes))

	c.JSON(http.StatusOK, *res)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, []map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, []map[string]interface{}{}, message, c)
}

// Fail 500
func Fail(c *gin.Context) {
	Result(ERROR, []map[string]interface{}{}, "fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, []map[string]interface{}{}, message, c)
}

func FailWithCodeMessage(code int, message string, c *gin.Context) {
	Result(code, []map[string]interface{}{}, message, c)
}
