package service

import "fmt"

/*
  @Author : zggong
*/

func IsLeader(username string) (status bool, err error) {

	resp, err := RequestGateway("https://eim.aibee.cn/api/eim-hr/v1/departments/", map[string]string{
		"at_time": "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resp.Data.([]interface{}) {
		if username == v.(map[string]interface{})["leader"].(string) {
			status = true
			return
		}
	}
	return
}
