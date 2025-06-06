package serializer

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Response Serializer
// All service should return HTTP 200，and `status` will tell the specific service status：
// status 0: service success
// status >0: service error
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`             // user message
	Error  string      `json:"error,omitempty"` // development error message
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

const (
	//1*** Client Error
	// StatusCheckLogin 未登录
	StatusCheckLogin = 1001
	// StatusNoRightErr 未授权访问
	StatusNoRightErr = 1002

	//2*** server Error
	// StatusDBError 数据库操作失败
	StatusDBError = 2001
	// StatusEncryptError 加密失败
	StatusEncryptError = 2002

	// StatusParamErr 参数错误
	StatusParamErr = 3001
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Status: StatusCheckLogin,
		Msg:    "User not logged in",
	}
}

// Success 成功响应
func Success(data interface{}) Response {
	return Response{
		Status: 0,
		Data:   data,
		Msg:    "success",
	}
}

// Err 通用错误处理
func Err(errStatus int, msg string, err error) Response {
	res := Response{
		Status: errStatus,
		Msg:    msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = fmt.Sprintf("%+v", err)
	}
	return res
}

func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "Database operation failed"
	}
	return Err(StatusDBError, msg, err)
}

func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "Parameter error"
	}
	return Err(StatusParamErr, msg, err)
}
