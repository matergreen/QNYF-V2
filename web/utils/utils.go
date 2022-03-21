package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"os"
)

// md5 加密
func EzMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// 写文件
func WritetoFile(s string) {
	fileobj, _ := os.OpenFile("uidList", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer fileobj.Close()

	w := bufio.NewWriter(fileobj)
	n, err := w.WriteString(s)
	if n != 0 && err == nil {
		w.Flush()
	}
}
