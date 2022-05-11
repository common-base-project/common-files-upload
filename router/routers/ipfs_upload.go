package routers

import (
	"fmt"
	"unicorn-files/handler/web3"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// IPFSUploadRouter 文件上传和下载
func IPFSUploadRouter(g *gin.Engine) {
	group := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/web3")
	uploadGroups := g.Group(group)
	{
		// Upload
		uploadGroups.POST("/ipfs", web3.UploadMutiFileHandler)
		uploadGroups.GET("/ipfs", web3.GetFileUrlByCidHandler)
	}
}
