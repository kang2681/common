package stringsext

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// MD5 计算字符串的 MD5 散列值
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Rand 字符串随机
func Rand(str string, num int) string {
	data := []rune(str)
	length := len(data)
	rs := make([]rune, 0, num)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		rs = append(rs, data[r.Intn(length)])
	}
	return string(rs)
}
