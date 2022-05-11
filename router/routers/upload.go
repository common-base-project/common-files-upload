package routers

import (
	"fmt"
	"unicorn-files/handler/upload"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// UploadRouter 文件上传
func UploadRouter(g *gin.Engine) {
	UploadRouterGroup := fmt.Sprintf("%s%s", viper.GetString(`api.version`), "/upload")
	UploadGroups := g.Group(UploadRouterGroup)
	{
		// Upload
		UploadGroups.POST("", upload.UploadMutiFileHandler)
	}
}
