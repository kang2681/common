package validate

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/asaskevich/govalidator"
)

func IsEmail(str string) bool {
	return govalidator.IsEmail(str)
}

//身份证格式校验
func CheckIdCard(idCard string) error {
	flag, err := regexp.MatchString(`^[\d]{17}[\dX]{1}$`, idCard)
	if err != nil {
		return err
	}
	if !flag {
		return fmt.Errorf("请输入正确的18位身份证号码")
	}
	idCardByte := []byte(idCard)
	id := idCardByte[:17]
	var sum int
	wi := [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2} //将前面的身份证号码17位数分别乘以不同的系数。从第一位到第十七位的系数分别为：7－9－10－5－8－4－2－1－6－3－7－9－10－5－8－4－2
	for index, value := range id {
		tmp, err := strconv.Atoi(string(value))
		if err != nil {
			return err
		}
		sum += tmp * wi[index] //将这17位数字和系数相乘的结果相加
	}
	a18 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2} //身份证格式固定最后一位
	var last byte = (idCardByte[17] - 48)
	if idCardByte[17] == 88 {
		last = 'X'
	}
	if a18[sum%11] == last {
		return nil
	}
	return fmt.Errorf("verify error")
}

//IsMobilePhoneNum 判断是不是手机号  支持165，166，198，199号段
func IsMobilePhoneNum(s string) (isRight bool) {
	phoneExpArr := []string{`^[1][34578][0-9]{9}$`, `^[1][6][5-6][0-9]{8}$`, `^[1][9][189][0-9]{8}$`}
	for _, pnoneExp := range phoneExpArr {
		reg := regexp.MustCompile(pnoneExp)
		if len(reg.FindAllString(s, -1)) > 0 {
			isRight = true
		}
	}
	return
}

//IsTelPhoneNum 判断是不是电话号码
func IsTelPhoneNum(s string) (isRight bool) {
	reg := regexp.MustCompile(`^0[\d]{2,3}-[\d]{7,8}$`)
	if len(reg.FindAllString(s, -1)) > 0 {
		isRight = true
	}
	return
}

//检查是否含有中文 是true 否false
func IsContainChinese(s string) bool {
	rs, _ := regexp.MatchString("[\u4e00-\u9fa5]+", s)
	return rs
}
