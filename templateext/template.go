package templateext

import (
	"bytes"
	"text/template"
)

func Parsef(text string, data interface{}, funcMap template.FuncMap) (string, error) {
	//  文件名解析
	tmpl, err := template.New("name").Parse(text)
	if err != nil {
		return "", err
	}
	if funcMap != nil && len(funcMap) > 0 {
		tmpl = tmpl.Funcs(funcMap)
	}
	buf := bytes.NewBuffer([]byte{})
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
