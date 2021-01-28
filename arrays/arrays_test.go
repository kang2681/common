package arrays

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInStringArray(t *testing.T) {
	Convey("TestInStringArray", t, func() {
		Convey("false", func() {
			flag := InStringArray("1", []string{"2", "3"})
			So(flag, ShouldBeFalse)
		})
		Convey("true", func() {
			flag := InStringArray("1", []string{"1", "2", "3"})
			So(flag, ShouldBeTrue)
		})
	})
}

func TestStringArrayUniqueMerge(t *testing.T) {
	Convey("TestStringArrayUniqueMerge", t, func() {
		Convey("不重复", func() {
			data := StringArrayUniqueMerge([]string{"1"}, []string{"2", "3"})
			So(data, ShouldResemble, []string{"1", "2", "3"})
		})
		Convey("重复", func() {
			data := StringArrayUniqueMerge([]string{"1", "2"}, []string{"2", "3"})
			So(data, ShouldResemble, []string{"1", "2", "3"})
		})
	})
}

func TestStringArrayUnique(t *testing.T) {
	Convey("TestStringArrayUnique", t, func() {
		Convey("不重复", func() {
			data := StringArrayUnique([]string{"2", "3"})
			So(data, ShouldResemble, []string{"2", "3"})
		})
		Convey("重复", func() {
			data := StringArrayUnique([]string{"1", "2", "1"})
			So(data, ShouldResemble, []string{"1", "2"})
		})
	})
}

func TestStringArrayIntersect(t *testing.T) {
	Convey("TestStringArrayIntersect", t, func() {
		Convey("不重复", func() {
			data := StringArrayIntersect([]string{"2", "3"}, []string{"1", "4"})
			So(data, ShouldResemble, []string{})
		})
		Convey("重复", func() {
			data := StringArrayIntersect([]string{"3", "2", "1"}, []string{"1", "2", "4"})
			So(data, ShouldResemble, []string{"2", "1"})
		})
	})
}

func TestStringArrayDiff(t *testing.T) {
	Convey("TestStringArrayDiff", t, func() {
		Convey("不重复", func() {
			data := StringArrayDiff([]string{"2", "3"}, []string{"1", "4"})
			So(data, ShouldResemble, []string{"2", "3"})
		})
		Convey("重复", func() {
			data := StringArrayDiff([]string{"3", "2", "1"}, []string{"1", "2", "4"})
			So(data, ShouldResemble, []string{"3"})
		})
	})
}

func TestStringArrayChunk(t *testing.T) {
	Convey("TestStringArrayChunk", t, func() {
		Convey("panic", func() {
			So(func() { StringArrayChunk([]string{"2", "3"}, 0) }, ShouldPanic)
		})
		Convey("success size 2", func() {
			data := StringArrayChunk([]string{"3", "2", "1"}, 2)
			So(data, ShouldResemble, [][]string{[]string{"3", "2"}, []string{"1"}})
		})
	})
}

func TestStringArrayCombine(t *testing.T) {
	Convey("TestStringArrayCombine", t, func() {
		Convey("error", func() {
			rs, err := StringArrayCombine([]string{"k1", "k2"}, []string{"v1"})
			So(err, ShouldBeError)
			So(rs, ShouldBeEmpty)
		})
		Convey("success", func() {
			rs, err := StringArrayCombine([]string{"k1", "k2"}, []string{"v1", "v2"})
			So(err, ShouldBeNil)
			So(rs, ShouldResemble, map[string]string{"k1": "v1", "k2": "v2"})
		})
		Convey("success empty", func() {
			rs, err := StringArrayCombine([]string{}, []string{})
			So(err, ShouldBeNil)
			So(rs, ShouldResemble, map[string]string{})
		})
	})
}

func TestStringArrayCountValues(t *testing.T) {
	Convey("TestStringArrayCountValues", t, func() {
		Convey("all diff", func() {
			rs := StringArrayCountValues([]string{"k1", "k2"})
			So(rs, ShouldResemble, map[string]int{"k1": 1, "k2": 1})
		})
		Convey("has same", func() {
			rs := StringArrayCountValues([]string{"k1", "k2", "k2"})
			So(rs, ShouldResemble, map[string]int{"k1": 1, "k2": 2})
		})

	})
}

func TestStringArraySearch(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			index, flag := StringArraySearch("1", []string{"k1", "k2", "k3"})
			So(flag, ShouldBeFalse)
			So(index, ShouldEqual, 0)
		})
		Convey("find", func() {
			index, flag := StringArraySearch("k1", []string{"k1", "k2", "k2"})
			So(flag, ShouldBeTrue)
			So(index, ShouldEqual, 0)
		})

	})
}
