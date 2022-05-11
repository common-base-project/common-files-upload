package models

/*
  @Author : zggong
*/

import (
	"unicorn-files/models/files"
	"unicorn-files/pkg/connection"
)

func AutoMigrateTable() {
	if connection.DbEnable {
		connection.DB.Self.AutoMigrate(
			&files.File{},
		)
	}
}
