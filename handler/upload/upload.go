package upload

import (
	"fmt"
	"strconv"
	"strings"
	"unicorn-files/models/files"
	"unicorn-files/pkg/logger"
	"unicorn-files/pkg/response/code"
	. "unicorn-files/pkg/response/response"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// UploadMutiFileHandler 上传文件接口
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
	logger.Info(isMulti)

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
		fileId := fmt.Sprintf("%s|%s", uuidString, file.Filename)
		err := c.SaveUploadedFile(file, fmt.Sprintf("%s%v", viper.GetString("filePath"), fileId))
		if err != nil {
			logger.Error(err)
			continue
		}

		// 存库
		data, err := saveFileToDB(fileId, file.Filename)
		if err != nil {
			logger.Error(err.Error())
		}

		fileDataList = append(fileDataList, data)
	}

	if len(fileDataList) <= 0 {
		logger.Errorf("上传文件失败: fileDataList %d", len(fileDataList))
		Response(c, code.UploadError, nil, "上传文件失败")
		return
	}

	if isMultipart {
		Response(c, nil, fileDataList, "上传成功")
	} else {
		Response(c, nil, fileDataList[0], "上传成功")
	}
}

// @saveFileToDB 保存文件到数据库
func saveFileToDB(fileId string, fileName string) (files.File, error) {
	data := files.File{
		FileId:   fileId,
		Url:      viper.GetString("uploadUrl") + fileId,
		FileName: fileName,
	}
	if err := data.Create(); err != nil {
		logger.Error(err.Error())
		return data, err
	}

	return data, nil
}
