package result

/*
  @Author : zggong
*/

import (
	"fmt"
	"net/http"
	"unicorn-files/pkg/logger"
	httpCode "unicorn-files/pkg/response/code"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	//log.Println(msg)
	// 开始时间
	c.JSON(http.StatusOK, ResponseData{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func Response(c *gin.Context, err error, data interface{}, resultText string) {
	code, message := httpCode.DecodeErr(err)

	if err != nil {
		message = fmt.Sprintf("%s，错误：%v", message, resultText)
	}

	if err == nil && resultText != "" {
		message = resultText
	}

	// write log
	if code != httpCode.Success.ErrNo {
		logger.Error(message)
	}

	// always return http.StatusOK
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  message,
		Data: data,
	})
}
