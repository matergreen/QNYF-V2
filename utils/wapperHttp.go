package utils

import (
	"io"
)

const (
	USERAGENT string = "Mozilla/5.0 Windows NT 10 Chrome/90.0.4832"
)

// 对 HTTP 方法进行简单封装
func WapperHttp(method, url string, data io.Reader) []byte {
	var result []byte
	switch method {
	case "GET":
		result = ezGet(url)
	case "POST":
		result = ezPost(url, data)
	}
	return result
}
