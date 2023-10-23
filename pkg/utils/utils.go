package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 加盐字符串
const secret = "哎呦你干嘛"

func Md5Sum(b []byte) string {
	h := md5.New()
	h.Write(b)
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(nil))
}
