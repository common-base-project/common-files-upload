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
	Errno  int         `json:"code"`
	Errmsg string      `json:"message"`
	Data   interface{} `json:"data"`
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
		Errno:  code,
		Errmsg: message,
		Data:   data,
	})
}
