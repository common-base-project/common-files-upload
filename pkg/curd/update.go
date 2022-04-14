package curd

import (
	"unicorn-files/pkg/connection"
)

/*
  @Author : zggong
*/

func Update(p *Param) (err error) {
	err = whereDB(p)

	err = connection.DB.Self.Model(p.Param).Save(p.Param).Error
	if err != nil {
		return
	}

	return nil
}
