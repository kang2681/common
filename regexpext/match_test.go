package regexpext

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewParser(t *testing.T) {
	Convey("TestRuTestNewParserneLen", t, func() {
		Convey("New", func() {
			rs := NewParser("{{code}}_{{name}}-log")
			So(rs, ShouldResemble, &Parser{format: "{{code}}_{{name}}-log", re: regexp.MustCompile("^(?P<code>[^_]*)_(?P<name>[^-]*)-log")})
		})
	})
}

func TestParser_ParseString(t *testing.T) {
	Convey("TestParser_ParseString", t, func() {
		Convey("New", func() {
			rs := NewParser("{{code}}_{{name}}.log")
			data, err := rs.ParseString("111_test.log")
			So(err, ShouldBeNil)
			So(data, ShouldResemble, map[string]string{"code": "111", "name": "test"})
		})
	})
}
