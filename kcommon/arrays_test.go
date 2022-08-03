package kcommon

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInArray(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			flag := InArray("1", []string{"k1", "k2", "k3"})
			So(flag, ShouldBeFalse)
		})
		Convey("find", func() {
			flag := InArray("k1", []string{"k1", "k2", "k3"})
			So(flag, ShouldBeTrue)
		})

	})
}

func TestArrayUniqueMerge(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr := ArrayUniqueMerge([]string{"1"}, []string{"1", "2"}, []string{"2", "3"})
			So(arr, ShouldResemble, []string{"1", "2", "3"})
		})
		Convey("find", func() {
			arr := ArrayUniqueMerge([]string{"1"}, []string{"1"}, []string{"2", "3"})
			So(arr, ShouldResemble, []string{"1", "2", "3"})
		})
	})
}

func TestArrayUnique(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr := ArrayUnique([]string{"1", "2", "1", "3"})
			So(arr, ShouldResemble, []string{"1", "2", "3"})
		})
		Convey("find", func() {
			arr := ArrayUnique([]string{"1", "2", "3"})
			So(arr, ShouldResemble, []string{"1", "2", "3"})
		})
	})
}

func TestArrayIntersect(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr := ArrayIntersect([]string{"1", "2", "1", "3"}, []string{"2", "3"}, []string{"3"})
			So(arr, ShouldResemble, []string{"3"})
		})
		Convey("find", func() {
			arr := ArrayIntersect([]string{"1", "2", "3"}, []string{"4"}, []string{"5"})
			So(arr, ShouldResemble, []string{})
		})
	})
}

func TestArrayDiff(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr := ArrayDiff([]string{"1", "2", "1", "3"}, []string{"2", "3"}, []string{"3"})
			So(arr, ShouldResemble, []string{"1", "1"})
		})
		Convey("find", func() {
			arr := ArrayDiff([]string{"1", "2", "3"}, []string{"4"}, []string{"5"})
			So(arr, ShouldResemble, []string{"1", "2", "3"})
		})
	})
}

func TestArrayChunk(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr := ArrayChunk([]string{"1", "2", "1", "3"}, 3)
			So(arr, ShouldResemble, [][]string{[]string{"1", "2", "1"}, []string{"3"}})
		})
		Convey("find", func() {
			arr := ArrayChunk([]string{"1", "2", "3"}, 1)
			So(arr, ShouldResemble, [][]string{
				[]string{"1"},
				[]string{"2"},
				[]string{"3"},
			})
		})
	})
}

func TestArrayCombine(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr, err := ArrayCombine([]string{"1", "2"}, []string{"2", "1"})
			So(err, ShouldBeNil)
			So(arr, ShouldResemble, map[string]string{"1": "2", "2": "1"})
		})
		Convey("find", func() {
			arr, err := ArrayCombine([]string{"1", "2"}, []string{"2", "1", "3"})
			So(err, ShouldBeError)
			So(arr, ShouldBeNil)
		})
	})
}

func TestArrayCountValues(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			arr := ArrayCountValues([]string{"1", "2"})
			So(arr, ShouldResemble, map[string]int{"1": 1, "2": 1})
		})
		Convey("find", func() {
			arr := ArrayCountValues([]string{"1", "2", "1"})
			So(arr, ShouldResemble, map[string]int{"1": 2, "2": 1})
		})
	})
}

func TestArraySearch(t *testing.T) {
	Convey("TestStringArraySearch", t, func() {
		Convey("not find", func() {
			idx, flag := ArraySearch("1", []string{"1", "2"})
			So(flag, ShouldBeTrue)
			So(idx, ShouldEqual, 0)
		})
		Convey("find", func() {
			idx, flag := ArraySearch("4", []string{"1", "2", "1"})
			So(flag, ShouldBeFalse)
			So(idx, ShouldEqual, 0)
		})
	})
}
