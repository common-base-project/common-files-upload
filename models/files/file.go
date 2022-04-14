package files

import (
	"unicorn-files/models/base"
	"unicorn-files/pkg/connection"
	"unicorn-files/pkg/logger"
)

/*
  @Author : zggong
  文件对象
*/
type File struct {
	base.Model
	FileId   string `gorm:"column:fileId; not null; unique; type:varchar(128);" json:"fileId" form:"fileId"`
	Url      string `gorm:"column:url; not null; type:varchar(256);" json:"url" form:"url"`
	FileName string `gorm:"column:fileName; type:varchar(256);" json:"fileName" form:"fileName"`
	Memo     string `gorm:"column:memo; type:varchar(256);" json:"memo" form:"memo"`
}

func (File) TableName() string {
	return "t_files"
}

func (g *File) Create() (err error) {
	if err = connection.DB.Self.Save(g).Error; err != nil {
		logger.Error(err.Error())
	}
	return err
}
