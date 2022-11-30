/*
  @Author : zggong
*/

package main

import (
	"fmt"
	"os"
	"unicorn-files/models"
	"unicorn-files/pkg/config"
	_ "unicorn-files/pkg/config"
	"unicorn-files/pkg/connection"
	"unicorn-files/pkg/logger"
	"unicorn-files/pkg/service/auth_rsync"
	"unicorn-files/router"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("#########1", os.Getenv("ENV_SERVER_MODE"))
	config.EnvMode = os.Getenv("ENV_SERVER_MODE")
	config.Initial()
	fmt.Println("#########2", config.EnvMode)
	logger.Initial()
	connection.Initial()
}

// @title 文件统一上传API文档
// @version 0.0.1
// @contact.name Gong Zhigang
// @contact.email zggong@aibee.com
// http://localhost:9080/api/v1/swagger/index.html
func main() {
	// 同步用户和部门数据
	go auth_rsync.Main()

	// 同步数据结构
	models.AutoMigrateTable()

	g := gin.New()

	if config.EnvMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.EnvMode == "staging" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 加载路由
	router.Load(g)

	// 运行程序
	logger.Infof("程序启动：%s", viper.GetString(`server.port`))
	err := g.Run(viper.GetString(`server.port`))
	if err != nil {
		logger.Error("启动失败")
		panic(fmt.Sprintf("程序启动失败：%v", err))
	}

	defer connection.DB.Close()
}
