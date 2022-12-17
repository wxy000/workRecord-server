package common

import (
	"net/http"
	"server/globals"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseC 返回值带条数
type ResponseC struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func ResultC(code int, msg string, count int64, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, ResponseC{
		code,
		msg,
		count,
		data,
	})
}

// Ok 成功
func Ok(c *gin.Context) {
	Result(globals.SUCCESS, "操作成功", map[string]interface{}{}, c)
}

// OkWithMsg 成功-自定义msg
func OkWithMsg(message string, c *gin.Context) {
	Result(globals.SUCCESS, message, map[string]interface{}{}, c)
}

// OkWithData 成功-自定义返回数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(globals.SUCCESS, "操作成功", data, c)
}

// OkWithDetailed 成功-自定义msg和返回数据
func OkWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(globals.SUCCESS, message, data, c)
}

// Fail 失败
func Fail(c *gin.Context) {
	Result(globals.ERROR, "操作失败", map[string]interface{}{}, c)
}

// FailWithMsg 失败-自定义msg
func FailWithMsg(message string, c *gin.Context) {
	Result(globals.ERROR, message, map[string]interface{}{}, c)
}

// FailWithDetailed 失败-自定义msg和返回数据
func FailWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(globals.ERROR, message, data, c)
}

// FailLogin 登录失败
func FailLogin(c *gin.Context) {
	Result(globals.ERRORLOGIN, "失败", map[string]interface{}{}, c)
}

// FailLoginWithMsg 登录失败-自定义msg
func FailLoginWithMsg(message string, c *gin.Context) {
	Result(globals.ERRORLOGIN, message, map[string]interface{}{}, c)
}

// FailLoginWithDetailed 登录失败-自定义msg和返回数据
func FailLoginWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(globals.ERRORLOGIN, message, data, c)
}

// OkC 成功（带条数）
func OkC(count int64, c *gin.Context) {
	ResultC(globals.SUCCESS, "操作成功", count, map[string]interface{}{}, c)
}

// OkWithMsgC 成功（带条数）-自定义msg
func OkWithMsgC(count int64, message string, c *gin.Context) {
	ResultC(globals.SUCCESS, message, count, map[string]interface{}{}, c)
}

// OkWithDataC 成功（带条数）-自定义返回数据
func OkWithDataC(count int64, data interface{}, c *gin.Context) {
	ResultC(globals.SUCCESS, "操作成功", count, data, c)
}

// OkWithDetailedC 成功（带条数）-自定义msg和返回数据
func OkWithDetailedC(count int64, message string, data interface{}, c *gin.Context) {
	ResultC(globals.SUCCESS, message, count, data, c)
}

// FailC 失败（带条数）
func FailC(count int64, c *gin.Context) {
	ResultC(globals.ERROR, "操作失败", count, map[string]interface{}{}, c)
}

// FailWithMsgC 失败（带条数）-自定义msg
func FailWithMsgC(count int64, message string, c *gin.Context) {
	ResultC(globals.ERROR, message, count, map[string]interface{}{}, c)
}

// FailWithDetailedC 失败（带条数）-自定义msg和返回数据
func FailWithDetailedC(count int64, message string, data interface{}, c *gin.Context) {
	ResultC(globals.ERROR, message, count, data, c)
}
