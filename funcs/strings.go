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

// Substr 返回字符串的子串
// str string
// 输入字符串。必须至少有一个字符。
// start
// 如果 start 是非负数，返回的字符串将从 string 的 start 位置开始，从 0 开始计算。例如，在字符串 “abcdef” 中，在位置 0 的字符是 “a”，位置 2 的字符串是 “c” 等等。
// 如果 start 是负数，返回的字符串将从 string 结尾处向前数第 start 个字符开始。
// 如果 string 的长度小于 start，将返回 false。
// length
// 如果提供了正数的 length，返回的字符串将从 start 处开始最多包括 length 个字符（取决于 string 的长度）。
// 如果提供了负数的 length，那么 string 末尾处的 length 个字符将会被省略（若 start 是负数则从字符串尾部算起）。如果 start 不在这段文本中，那么将返回 false。
// 如果提供了值为 0，false 或 null 的 length，那么将返回一个空字符串。
func Substr(str string, start, length int) (string, bool) {
	data := []rune(str)
	strLen := len(data)
	var begin, end int
	if start < 0 {
		begin = strLen + start
		if begin < 0 {
			begin = 0
		}
	} else {
		begin = start
	}
	if length < 0 {
		end = strLen + length
	} else {
		end = start + length
	}
	if end < 0 || begin >= strLen || begin > end {
		return "", false
	}
	if end > strLen {
		end = strLen
	}
	return string(data[begin:end]), true
}

func SubStrToEnd(str string, start int) (string, bool) {
	data := []rune(str)
	strLen := len(data)
	if start > strLen {
		return "", false
	}
	var begin int = start
	if start < 0 {
		begin = strLen + start
		if begin < 0 {
			begin = 0
		}
	}
	if begin < 0 {
		begin = 0
	}
	if begin > strLen {
		return "", false
	}
	return string(data[begin:strLen]), true
}
