package vo

import (
	"github.com/gin-gonic/gin"
	"time"
	biz_err "yinglian.com/yun-ai-server/internal/error"
	"yinglian.com/yun-ai-server/internal/request"
)

type Result struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"requestId"`
	TimeStamp int64       `json:"timeStamp"`
}

func Success(data interface{}, msg string, c *gin.Context) Result {
	return Result{
		Code:      0, // 0 表示成功
		Msg:       msg,
		Data:      data,
		RequestId: request.GetRequestID(c),
		TimeStamp: time.Now().UnixNano(),
	}
}
func Fail(err *biz_err.Err, data interface{}, c *gin.Context) Result {
	return Result{
		Code:      err.Code,
		Msg:       err.Msg,
		Data:      data,
		RequestId: request.GetRequestID(c),
		TimeStamp: time.Now().UnixNano(),
	}
}
