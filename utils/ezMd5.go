package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密
func EzMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
