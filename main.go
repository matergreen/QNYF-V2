package main

import (
	"QnyfV2/api"
	"QnyfV2/config"
	"QnyfV2/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

// 定义接受内容的结构体
type ImgInfo struct {
	Key string `json:"key"`
	Img string `json:"img"`
}

type Verify struct {
	Code int    `json:"code"`
	Info string `info:"info"`
	Data ImgInfo
}

// 定义 SAVEPATH 发送的 Json 结构体
type SaveInfo struct {
	UID      int
	UserType int
	JWD      string // 经纬度
	Key      string `json:"key"`
	Code     string `json:"code"` // 验证码
	ZZDKID   int
	A1       string // 正常
	A4       string // 无
	A2       string // 全部正常
	A3       string // 地址
	A11      string // 是否在校
	A12      string // 是否实习
	A13      string // 所处地区风险程度
	YXDM     string // 院校代码
	Version  string // 版本号
}

func main() {
	var idAndUids []string
	var key, code, msg string

	// 从文件读取 uids
	idAndUids = utils.UniqueSlice(utils.ReadFile())

	// 防止 死循环
	depth := 0
	for _, idAndUid := range idAndUids {
		// 2. 请求 IsClockIn
		idUidSlice := strings.Split(idAndUid, ",")
		fmt.Println(idUidSlice)
		id, uidstr := idUidSlice[0], idUidSlice[1]

		time.Sleep(3 * time.Second)
		urlChockIn := fmt.Sprintf(config.CheckIn+"?uid=%s&usertype=1&yxdm=10623", uidstr)
		contentChock := utils.WapperHttp("GET", urlChockIn, nil)

		// 判断今日是否打卡
		jsonObjClock, _ := simplejson.NewJson(contentChock)
		if jsonObjClock.Get("code").MustInt() == 400 {
			tmpmsg := fmt.Sprintf("%s 用户接口非法访问...\n", id)
			msg += tmpmsg
			continue
		}
		if jsonObjClock.GetPath("data", "msg").MustString() != "未打卡" {
			tmpmsg := fmt.Sprintf("%s 今日已打卡!\n", id)
			msg += tmpmsg
			continue
		}

		for {
			// 3. 获取验证码
			time.Sleep(1 * time.Second)
			key, code = GetVerify()
			if key == code {
				tmpmsg := fmt.Sprintf("%s 今日打卡失败!\n", id)
				msg += tmpmsg
				continue
			}
			fmt.Println(key, code)

			// 4. 签到
			intUid, _ := strconv.Atoi(uidstr)
			ok := CheckIn(intUid, key, code)
			if ok {
				tmpmsg := fmt.Sprintf("%s 今日打卡成功!\n", id)
				msg += tmpmsg
				break
			}
			time.Sleep(1 * time.Second)
			if depth == 10 {
				tmpmsg := fmt.Sprintf("%s 今日打卡失败!\n", id)
				msg += tmpmsg
				break
			}
			depth++
		}
	}

	api.SendMsg(msg)
}

// 初始化发送信息结构体
func InitSaveInfo(uid int, jwd, key, code, a3, a11, a12 string) *SaveInfo {
	return &SaveInfo{
		UID:      uid,
		UserType: 1,
		JWD:      jwd,
		Key:      key,
		Code:     code,
		ZZDKID:   37,
		A1:       "正常",
		A4:       "无",
		A2:       "全部正常",
		A3:       a3,
		A11:      a11,
		A12:      a12,
		A13:      "地风险区",
		YXDM:     "10623",
		Version:  "v1.3.2",
	}
}

// 获取验证码
func GetVerify() (string, string) {
	var allResult Verify
	for {
		respData := utils.WapperHttp("GET", config.VERIFYCODEPATH, nil)
		_ = json.Unmarshal(respData, &allResult)

		imgInfo := allResult.Data
		utils.DecodetoImg(imgInfo.Img)
		code := api.GetCode()
		if code == "" {
			return "", ""
		}
		code = strings.Replace(code, " ", "", -1)
		if len(code) == 4 {
			return allResult.Data.Key, code
		}
		time.Sleep(1 * time.Second)
	}
}

// 打卡
func CheckIn(uid int, key, code string) bool {
	JWD := "123123,123123"
	A11 := "在校"
	A12 := "未实习"

	seninfo := InitSaveInfo(uid, JWD, key, code, config.A3, A11, A12)
	sendjson, _ := json.Marshal(seninfo)
	jsonCheckin := utils.WapperHttp("POST", config.SAVEPATH, bytes.NewReader(sendjson))
	result, _ := simplejson.NewJson(jsonCheckin)
	fmt.Println(result)

	return result.GetPath("code").MustInt() == 200
}
