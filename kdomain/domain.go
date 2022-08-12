package kdomain

import (
	"regexp"
	"strings"
)

// GetDomainTld 获取域名最后一级后缀
func GetDomainTld(domain string) string {
	dnArr := strings.Split(domain, ".")
	return strings.ToLower(dnArr[len(dnArr)-1])
}

// GetDomainSuffix 获取域名后缀
func GetDomainSuffix(domain string) string {
	arr := strings.Split(domain, ".")
	domainClassAll := ""
	arrSize := len(arr)
	if arrSize == 3 {
		domainClassAll = arr[1] + "." + arr[2]
	} else if arrSize == 2 {
		domainClassAll = arr[1]
	}
	return strings.ToLower(domainClassAll)
}

// GetDomainBody 获取域名前缀
func GetDomainBody(domain string) string {
	n := strings.Index(domain, ".")
	if n <= 0 {
		return ""
	}
	return strings.ToLower(domain[0:n])
}

// IsCnnicDomain 判断域名后缀是否是CN 中国 公司 网络 网址
func IsCnnicDomain(dn string) bool {
	tld := GetDomainTld(dn)
	if tld == "cn" || tld == "中国" || tld == "公司" || tld == "网络" || tld == "网址" {
		return true
	}
	return false
}

// IsDomainBodyHaveZhongwen 判断域名主体是否有中文
func IsDomainBodyHaveZhongwen(domain string) bool {
	domainBody := GetDomainBody(domain)
	reg := regexp.MustCompile(`[\p{Han}]`)
	if nil == reg.FindAllString(domainBody, -1) {
		return false
	}
	return true
}

// IsDomainSuffixHaveZhongwen 判断是否含有中文的后缀
func IsDomainSuffixHaveZhongwen(domain string) bool {
	suffix := GetDomainSuffix(domain)
	reg := regexp.MustCompile(`[\p{Han}]`)
	if nil == reg.FindAllString(suffix, -1) {
		return false
	}
	return true
}
