package funcs

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 计算字符串的 MD5 散列值
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//strlen — 获取字符串字符长度
func StrCharlen(str string) int {
	return len([]rune(str))
}
