package serializer

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// StatusCheckLogin 未登录
	StatusCheckLogin = 401
	// StatusNoRightErr 未授权访问
	StatusNoRightErr = 403
	// StatusDBError 数据库操作失败
	StatusDBError = 50001
	// StatusEncryptError 加密失败
	StatusEncryptError = 50002
	//StatusParamErr 各种奇奇怪怪的参数错误
	StatusParamErr = 40001
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Status: StatusCheckLogin,
		Msg:    "未登录",
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

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(StatusDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(StatusParamErr, msg, err)
}
