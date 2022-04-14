package upload

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicorn-files/models/files"
	"unicorn-files/pkg/connection"
	"unicorn-files/pkg/logger"
	"unicorn-files/pkg/response/code"
	. "unicorn-files/pkg/response/response"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

/*
  @Author : zggong
*/

// always return http.StatusOK
type UploadResponse struct {
	Code   int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// @Summary 上传文件接口
// @Description 上传文件接口
// @Tags 文件
// @Accept json
// @Produce json
// @Param isMultipart query string false "是否为多文件上传"
//// @Param upload_type query string false "上传文件类型 images/file/media"
// @Param file formData file true "文件集"
// @Success 0
// @Router /api/v1/upload [post]
func UploadMutiFileHandler(c *gin.Context) {
	isMulti := c.DefaultQuery("isMultipart", "false")
	//uploadType := c.DefaultQuery("upload_type", "images")
	logger.Info(isMulti)
	//logger.Info(uploadType)

	//var filesMax int64
	//if uploadType == "images" {
	//	filesMax = 500 << 20 // 2M
	//} else if uploadType == "file" {
	//	filesMax = 500 << 20 // 50M
	//} else if uploadType == "media" {
	//	filesMax = 5000 << 20 // 500M
	//}
	//
	//c.Request.Body = http.MaxBytesReader(nil, c.Request.Body, filesMax)
	//// 设置文件大小，文件最大为2M
	//err := c.Request.ParseMultipartForm(filesMax)
	//if err != nil {
	//	Response(c, code.InternalServerError, nil, "文件太大了")
	//	return
	//}
	// =============================================================
	us, _ := uuid.NewV4()
	uuidString := us.String()
	uuidString = strings.ReplaceAll(uuidString, "-", "")[0:10]

	form, err := c.MultipartForm()
	if form == nil {
		logger.Errorf("没有上传文件: %s", err.Error())
		Response(c, code.UploadError, nil, err.Error())
		return
	}
	formFiles := form.File["file"]
	isMultipart, _ := strconv.ParseBool(isMulti)

	var fileDataList []files.File
	for _, file := range formFiles {
		fileName := fmt.Sprintf("%s|%s", uuidString, file.Filename)
		err := c.SaveUploadedFile(file, fmt.Sprintf("%s%v", viper.GetString("filePath"), fileName))
		if err != nil {
			logger.Error(err)
			continue
		}

		data := files.File{
			FileId:   fileName,
			Url:      viper.GetString("uploadUrl") + fileName,
			FileName: file.Filename,
		}
		if err := data.Create(); err != nil {
			logger.Error(err.Error())
		}
		fileDataList = append(fileDataList, data)
	}

	if len(fileDataList) <= 0 {
		logger.Errorf("上传文件失败: fileDataList %d", len(fileDataList))
		Response(c, code.UploadError, nil, "上传文件失败")
		return
	}
	// 保存文件到db
	//go saveFile(fileDataList)

	if isMultipart {
		c.JSON(http.StatusOK, UploadResponse{
			Code:   0,
			Errmsg: "",
			Data:   fileDataList,
		})
	} else {
		c.JSON(http.StatusOK, UploadResponse{
			Code:   0,
			Errmsg: "",
			Data:   fileDataList[0],
		})
	}
}

// 保存文件
func saveFile(fileDataList []files.File) {
	for _, file := range fileDataList {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
			}
		}()

		if err := connection.DB.Self.Save(file).Error; err != nil {
			logger.Error(err.Error())
		}

	}
}
