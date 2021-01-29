package regexpext

import "regexp"

func ParseRegexp(s string) (*regexp.Regexp, error) {
	if s == "" {
		return nil, nil
	}
	reg, err := regexp.Compile(s)
	if err != nil {
		return nil, err
	}
	return reg, nil
}

// 排除
func IsExcluded(str string, excl *regexp.Regexp) bool {
	return excl != nil && excl.MatchString(str)
}

// 包含
func IsIncluded(str string, only *regexp.Regexp) bool {
	return only == nil || only.MatchString(str)
}
