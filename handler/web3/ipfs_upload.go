package web3

import (
	"context"
	"fmt"
	"github.com/web3-storage/go-w3s-client"
	"io"
	"os"
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

// UploadMutiFileHandler @Summary 上传文件接口
// @Description 上传文件接口
// @Tags 文件
// @Accept json
// @Produce json
// @Param isMultipart query string false "是否为多文件上传"
//// @Param upload_type query string false "上传文件类型 images/file/media"
// @Param file formData file true "文件集"
// @Success 0
// @Router /api/v1/web3/ipfs [post]
func UploadMutiFileHandler(c *gin.Context) {
	isMulti := c.DefaultQuery("isMultipart", "false")
	//uploadType := c.DefaultQuery("upload_type", "images")
	logger.Info(isMulti)
	//logger.Info(uploadType)
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

	web3Client, err := w3s.NewClient(w3s.WithToken(viper.GetString(`web3.token`)))
	if err != nil {
		logger.Error(err)
		Response(c, code.UploadError, nil, err.Error())
		return
	}

	var fileDataList []files.File
	for _, file := range formFiles {
		fileId := fmt.Sprintf("%s|%s", uuidString, file.Filename)
		dst := fmt.Sprintf("%s%v", viper.GetString("filePath"), fileId)
		// ==============================文件处理===============================
		src, err := file.Open()
		if err != nil {
			logger.Error(err)
			continue
		}
		defer src.Close()

		out, err := os.Create(dst)
		if err != nil {
			logger.Error(err)
			continue
		}
		defer out.Close()

		_, err = io.Copy(out, src)
		// =============================================================

		// Write a file/directory
		cid, err := web3Client.Put(context.Background(), out)
		if err != nil {
			logger.Error(err)
			continue
		}
		// fmt.Printf("https://ipfs.io/ipfs/%s\n", cid)
		url := fmt.Sprintf("https://%v.ipfs.dweb.link", cid)

		// 存库
		data, err := saveFileToDB(cid.String(), url, file.Filename, fileId)
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
	// 保存文件到db
	//go saveFile(fileDataList)

	if isMultipart {
		Response(c, code.Success, fileDataList, "上传成功")
	} else {
		Response(c, code.Success, fileDataList[0], "上传成功")
	}
}

// @saveFileToDB 保存文件到数据库
func saveFileToDB(fileId string, url string, fileName string, dst string) (files.File, error) {
	data := files.File{
		FileId:   fileId,
		Url:      url,
		FileName: fileName,
		Memo:     dst,
	}
	if err := data.Create(); err != nil {
		logger.Error(err.Error())
		return data, err
	}

	return data, nil
}

// GetFileUrlByCidHandler 查询文件接口
// @Description 查询文件接口
// @Tags 文件
// @Accept json
// @Produce json
// @Param cid query string true "文件集"
// @Success 0
// @Router /api/v1/web3/ipfs [get]
func GetFileUrlByCidHandler(c *gin.Context) {
	cid := c.DefaultQuery("cid", "")
	var data files.File

	if connection.DbEnable {
		err := connection.DB.Self.Model(&files.File{FileId: cid}).First(&data).Error
		if err != nil {
			logger.Error(err.Error())
			Response(c, code.UploadError, nil, "文件不存在")
			return
		}
	}

	Response(c, code.Success, data, "查询成功")
}

func DeleteFileByCidHandler(c *gin.Context) {
	cid := c.DefaultQuery("cid", "")

	// 带额外条件的删除
	//db.Where("name = ?", "jinzhu").Delete(&email)
	if connection.DbEnable {
		err := connection.DB.Self.Where("fileId = ?", cid).Delete(&files.File{}).Error
		if err != nil {
			logger.Error(err.Error())
			Response(c, code.UploadError, nil, "文件不存在")
			return
		}
	}

	Response(c, code.Success, cid, "删除成功")
}
