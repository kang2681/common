package kstring

import (
	"testing"

	"github.com/sirupsen/logrus"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMD5(t *testing.T) {
	Convey("TestMD5", t, func() {
		Convey("123456", func() {
			data := MD5("123456")
			So(data, ShouldEqual, "e10adc3949ba59abbe56e057f20f883e")
		})
		Convey("empty string", func() {
			data := MD5("")
			So(data, ShouldEqual, "d41d8cd98f00b204e9800998ecf8427e")
		})
	})
}

func TestRand(t *testing.T) {
	Convey("TestRand", t, func() {
		Convey("", func() {
			data := Rand("12", 2)
			logrus.Infof("Data:%s", data)
			So(data, ShouldBeIn, []string{"11", "22", "12", "21"})
		})
	})
}
