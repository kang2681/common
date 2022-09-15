package kregexp

import (
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
	format string
	re     *regexp.Regexp
}

func NewParser(format string) *Parser {
	reg := regexp.MustCompile(`\\\{\\\{([a-z_]+)\\\}\\\}(\\?(.))`)
	re := reg.ReplaceAllString(regexp.QuoteMeta(format+" "), "(?P<$1>[^$3]*)$2")
	return &Parser{
		format: format,
		re:     regexp.MustCompile(fmt.Sprintf("^%v", strings.Trim(re, " "))),
	}
}

func (p *Parser) ParseString(line string) (map[string]string, error) {
	re := p.re
	fields := re.FindStringSubmatch(line)
	if fields == nil {
		return nil, fmt.Errorf("line '%v' does not match given format '%v'", line, re)
	}
	data := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		data[name] = fields[i]
	}
	return data, nil
}
