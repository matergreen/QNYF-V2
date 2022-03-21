package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"qnyfweb/config"
	"qnyfweb/utils"

	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

// 登录表单结构体
type LoginInfo struct {
	YXDM     string // 院校代码
	UserType int    // 登录类型
	XGH      string // 学号
	Name     string // 姓名
	PassWord string // 密码
}

func main() {
	r := gin.Default()
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		t, _ := template.ParseFiles("index.html")
		_ = t.Execute(c.Writer, "")
	})

	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		schoolId := ctx.PostForm("schoolid")
		password := ctx.PostForm("password")

		// 1. 登录
		// 初始化登录结构体,并转换为 Json 格式
		loginInfo := InitLoginInfo(username, schoolId, password)
		loginJson, _ := json.Marshal(loginInfo)

		// 提交登录信息,如果登录成功,获取 uid ,否则程序退出
		resp, _ := http.Post(config.LOGIN, "application/json;charset=utf-8", bytes.NewReader(loginJson))

		reader, _ := ioutil.ReadAll(resp.Body)
		loginResp, _ := simplejson.NewJson(reader)
		name := loginResp.GetPath("data", "XM").MustString()
		uid := loginResp.GetPath("data", "ID").MustInt()
		saveLine := fmt.Sprintf("%s,%d\n", name, uid)

		var msg interface{}
		if loginResp.Get("code").MustInt() == 200 {
			msg = struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{200, "登录成功"}
			utils.WritetoFile(saveLine)
		} else {
			msg = loginResp
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	r.Run("0.0.0.0:8077")
}

// 初始化登录信息结构提
func InitLoginInfo(name, xgh, passwd string) *LoginInfo {
	return &LoginInfo{
		YXDM:     "10623",
		UserType: 1,
		XGH:      xgh,
		Name:     name,
		PassWord: utils.EzMD5(passwd),
	}
}
