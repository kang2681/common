package stringsext

import (
	"fmt"
)

// RuneLen rune字符串长度计算
func RuneLen(str string) int {
	return len([]rune(str))
}

// RuneSubstr rune字符串截取
// start > 0 从 start 位置开始， 如果start 大于字符串长度 err 不为空
// start < 0 从 尾部倒回计算， 如果 start+ 字符串长度 小于  0 将从 0 开始计算
// length > 0 从start 开始向后计算长度
// length < 0 从字符串末尾往前计算结束位置
func RuneSubstr(str string, start, length int) (string, error) {
	data := []rune(str)
	strLen := len(data)
	var begin int // 起始位置计算
	if start < 0 {
		begin = strLen + start
	} else {
		begin = start
	}
	if begin > strLen {
		return "", fmt.Errorf("start position greater than string length")
	}
	if begin < 0 {
		begin = 0
	}
	var end int // 结束位置计算
	if length < 0 {
		end = strLen + length
	} else {
		end = begin + length
	}
	// 综合判断
	if end < 0 {
		return "", fmt.Errorf("end position less than zero")
	}
	if begin > end {
		return "", fmt.Errorf("end position greater than start position")
	}
	if end > strLen {
		end = strLen
	}
	return string(data[begin:end]), nil
}

// RuneSubstrToEnd rune字符串截取尾部
// start > 0 从 start 位置开始， 如果start 大于字符串长度 err 不为空
// start < 0 从 尾部倒回计算， 如果 start+ 字符串长度 小于  0 将从 0 开始计算
func RuneSubstrToEnd(str string, start int) (string, error) {
	data := []rune(str)
	strLen := len(data)
	var begin int
	if start < 0 {
		begin = start + strLen
	} else {
		begin = start
	}
	if begin > strLen {
		return "", fmt.Errorf("start position greater than string length")
	}
	if begin < 0 {
		begin = 0
	}
	return string(data[begin:strLen]), nil
}
