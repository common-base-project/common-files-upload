package router

/*
  @Author : zggong
*/

import (
	"fmt"
	"net/http"
	"time"
	_ "unicorn-files/docs"
	"unicorn-files/pkg/logger"
	result "unicorn-files/pkg/response/response"
	"unicorn-files/pkg/tools"
	"unicorn-files/pkg/utils"
	"unicorn-files/router/routers"

	"github.com/spf13/viper"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 加载路由
func Load(g *gin.Engine) {
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, result.ResponseData{
			Errno:  404,
			Errmsg: "API地址不存在",
			Data:   nil,
		})
	})

	// pprof router
	pprof.Register(g)

	//cors， 跨域
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	// 注册zap相关中间件
	g.Use(cors.New(config))
	//g.Use(logger.GinLogger(), logger.GinRecovery(true))
	g.Use(utils.CostTime())
	// 开启中间件 允许使用跨域请求
	g.Use(Cors())

	// ========================文件配置===============================
	filePath := viper.GetString("filePath")
	_, err := tools.CreateDictByPath(filePath)
	if err != nil {
		logger.Error("创建目录失败，请手动创建![%v]\n", err)
		return
	}
	logger.Infof("创建目录成功: %s", filePath)

	staticPath := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/upload")
	//静态文件地址 http://localhost:port/api/v1/upload/fileid.jpg
	g.Static(staticPath, filePath)

	// g.POST("/api/v1/upload", upload.UploadMutiFileHandler)
	// 设置文件大小，文件最大为10M (默认 32 MiB)	32 << 20
	g.MaxMultipartMemory = 500 << 20 // 500M
	// =======================================================

	// swagger api docs
	g.GET(fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/swagger/*any"), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// jwt 检查
	//g.Use(middleware.CheckToken())

	// upload
	routers.UploadRouter(g)
	// ipfs 文件上传和下载
	routers.IPFSUploadRouter(g)

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
