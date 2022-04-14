package service

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
)

/*
  @Author : zggong
*/

type MyResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type AccessToken struct {
	RequestTime int64    `json:"request_time"`
	User        string   `json:"user"`
	Groups      []string `json:"groups"`
	ServiceName string   `json:"service_name"`
}

var secret interface{} = []byte("D&023u@981jwoIie_!@#*s;lij!poW2ireJLAn3)-")

const TokenNameInHeader = "Access-Token"

// 生成token返回tokenString用于设置http header
func (at *AccessToken) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user":         at.User,
		"request_time": at.RequestTime,
		"groups":       at.Groups,
		"service_name": at.ServiceName,
	})
	if tokenString, err := token.SignedString(secret); err != nil {
		//common.Log.Errorf("Token signaure failed: %s", err.Error())
		return ""
	} else {
		return tokenString
	}
}

// 请求网关数据
func RequestGateway(URL string, params map[string]string) (cResp MyResp, err error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	//var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	token := AccessToken{ServiceName: "case", RequestTime: time.Now().Unix()}
	TokenValue := token.GenerateToken()
	req := client.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		SetHeader(TokenNameInHeader, TokenValue)
	resp, err := req.Get(URL)

	if err != nil {
		log.Println("获取数据失败!", err)
		return
	}
	err = json.Unmarshal(resp.Body(), &cResp)
	if err != nil {
		log.Printf("JSON解析失败-%s", err.Error())
		return
	}
	if cResp.Code != 100000 {
		log.Printf("获取数据失败-%d-%s", cResp.Code, cResp.Msg)
		return
	}

	return
}
