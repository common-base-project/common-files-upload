package curd

import (
	"fmt"
	"unicorn-files/pkg/connection"
)

/*
  @Author : zggong
*/
type Param struct {
	Name       string
	Models     interface{}
	Param      interface{}
	WhereMap   map[string]interface{}
	WhereValue string
}

func whereDB(p *Param) (err error) {
	if !connection.DbEnable {
		return nil
	}
	db := connection.DB.Self
	var dataCount int64 = 0
	if p.WhereValue != "" {
		db = db.Where("name = ?", p.WhereValue)
	}

	if p.WhereMap != nil {
		for key, value := range p.WhereMap {
			db = db.Where(fmt.Sprintf("%v = ?", key), value)
		}
	}

	err = db.Model(p.Models).Count(&dataCount).Error
	if err != nil {
		err = fmt.Errorf("查询%s数据失败，%v", p.Name, err)
		return
	}

	if dataCount > 0 {
		err = fmt.Errorf("`%s`数据筛选出现问题，请确认", p.Name)
		return
	}

	return
}

func Create(p *Param) (err error) {
	if !connection.DbEnable {
		return nil
	}
	err = whereDB(p)
	if err != nil {
		return
	}

	err = connection.DB.Self.Save(p.Param).Error
	if err != nil {
		return
	}

	return nil
}
