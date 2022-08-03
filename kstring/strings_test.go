package kstring

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRuneLen(t *testing.T) {
	Convey("TestRuneLen", t, func() {
		Convey("A-Z0-9!@#$%^&*", func() {
			length := RuneLen("qwertyuiopasdfghjklzxcvbnm!@#$%^^&123456789*")
			So(length, ShouldEqual, 44)
		})
		Convey("中文", func() {
			length := RuneLen("你猜我猜不猜")
			So(length, ShouldEqual, 6)
		})
	})
}

func TestRuneSubstr(t *testing.T) {
	Convey("TestRuneSubstr", t, func() {
		Convey("0 ~ -1", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", 0, -1)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "你猜我猜不猜abcde")
		})
		Convey("0 ~ 0", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", 0, 0)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "")
		})
		Convey("-8 ~ -2", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", -8, -2)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "不猜abcd")
		})
		Convey("-2 ~ -8", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", -2, -8)
			So(err, ShouldBeError)
			So(substr, ShouldEqual, "")
		})
		Convey("-2 ~ 8", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", -2, 8)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "ef")
		})
		Convey("-2 ~ 0", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", -2, 0)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "")
		})
		Convey("-20 ~ 0", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", -20, 0)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "")
		})
		Convey("20 ~ 0", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", 20, 0)
			So(err, ShouldBeError)
			So(substr, ShouldEqual, "")
		})
		Convey("1 ~ -20", func() {
			substr, err := RuneSubstr("你猜我猜不猜abcdef", 1, -20)
			So(err, ShouldBeError)
			So(substr, ShouldEqual, "")
		})
	})
}

func TestRuneSubstrToEnd(t *testing.T) {
	Convey("TestRuneSubstrToEnd", t, func() {
		Convey("0", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", 0)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "你猜我猜不猜abcdef")
		})
		Convey("20", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", 20)
			So(err, ShouldBeError)
			So(substr, ShouldEqual, "")
		})
		Convey("-20", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", -20)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "你猜我猜不猜abcdef")
		})
		Convey("-12", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", -12)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "你猜我猜不猜abcdef")
		})
		Convey("12", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", 12)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "")
		})
		Convey("13", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", 13)
			So(err, ShouldBeError)
			So(substr, ShouldEqual, "")
		})
		Convey("-1", func() {
			substr, err := RuneSubstrToEnd("你猜我猜不猜abcdef", -1)
			So(err, ShouldBeNil)
			So(substr, ShouldEqual, "f")
		})
	})
}
