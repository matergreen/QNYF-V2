package utils

import (
	"encoding/base64"
	"io/ioutil"
)

func DecodetoImg(s string) {
	verifyCodeImg, _ := base64.StdEncoding.DecodeString(s)
	_ = ioutil.WriteFile("./verifycode.png", verifyCodeImg, 0666)
}
