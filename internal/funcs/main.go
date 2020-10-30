package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
)

var arr = []string{"string", "int8", "int16", "int32", "rune", "int", "int64", "uint8", "byte", "uint16", "uint32", "uint", "uint64", "float32", "float64"}

type Param struct {
	Name     string
	FuncName string
}

func main() {
	name := "arrays.tpl"
	t, err := template.New(name).ParseFiles(name)
	if err != nil {
		logrus.Errorf("template new error:%s", err.Error())
		return
	}
	data := bytes.NewBuffer([]byte{})
	vals := []Param{}
	for _, v := range arr {
		tmp := Param{Name: v, FuncName: strings.Title(v)}
		vals = append(vals, tmp)
	}
	if err := t.Execute(data, map[string]interface{}{"data": vals}); err != nil {
		logrus.Errorf("execute error:%s", err.Error())
		return
	}
	rs, err := format.Source(data.Bytes())
	if err != nil {
		logrus.Errorf("format source error:%s", err.Error())
		return
	}
	ioutil.WriteFile("./test.go", rs, os.ModePerm)
	//fmt.Printf("%s", rs)
}
