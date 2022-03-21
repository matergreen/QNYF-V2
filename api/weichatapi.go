package api

import (
	"QnyfV2/config"
	"QnyfV2/utils"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

// 模板参数结构体
type Data struct {
	Content struct {
		Value string `json:"value"`
	} `json:"content"`
}

// 模板请求结构体
type TemplateData struct {
	Touser      string `json:"touser"`
	Template_id string `json:"template_id"`
	Data        Data   `json:"data"`
}

// 获取 access_token
func getWCTken() string {

	contentByte := utils.WapperHttp("GET", fmt.Sprint(config.WECHATTOKENURL), nil)
	contentJson, _ := simplejson.NewJson(contentByte)

	return contentJson.GetPath("access_token").MustString()
}

// 调用模板 发送信息
func SendMsg(s string) {
	datas := TemplateData{
		Touser:      config.WECHATID,
		Template_id: config.TEMPLATEID,
		Data: Data{
			Content: struct {
				Value string "json:\"value\""
			}{Value: s},
		},
	}

	postData, _ := json.Marshal(datas)
	_ = utils.WapperHttp("POST", config.WECHATSENDMSGURL+getWCTken(), bytes.NewReader(postData))
}
