package api

import (
	"QnyfV2/config"
	"QnyfV2/utils"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/bitly/go-simplejson"
)

// 获取 token
func getToken() string {
	for {
		resp := utils.WapperHttp("GET", config.OAUTH, nil)
		tokenResp, _ := simplejson.NewJson(resp)
		accessToken, err := tokenResp.Get("access_token").String()
		if err == nil {
			return accessToken
		}
	}
}

// 获取验证码
func GetCode() string {
	token := getToken()

	uri, err := url.Parse(config.APIPATH)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", token)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile(config.VERIFYIMGPATH)
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendBody.Form.Add("language_type", "CHN_ENG")
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// 解析相应中的 code
	codeResp, _ := simplejson.NewJson(result)
	codeJson := codeResp.Get("words_result").MustArray()
	for _, k := range codeJson {
		v := k.(map[string]interface{})
		for _, code := range v {
			return code.(string)
		}
	}
	return ""
}
