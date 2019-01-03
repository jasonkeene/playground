package search

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearchCorrectness(t *testing.T) {
	Convey("Given an ordered slice of integers", t, func() {
		data := []int{0, 1, 3, 6, 9, 42, 53, 71, 90, 567, 998, 1000, 10001}
		Convey("it should find members", func() {
			So(BinarySearch(data, 90), ShouldEqual, 8)
		})
		Convey("it should return -1 if not found", func() {
			So(BinarySearch(data, 999), ShouldEqual, -1)
		})
		Convey("it should return -1 if value is smaller than 0th element", func() {
			So(BinarySearch(data, -20), ShouldEqual, -1)
		})
		Convey("it should return -1 if value is larger than nth element", func() {
			So(BinarySearch(data, 99999), ShouldEqual, -1)
		})
	})

	Convey("Given an ordered two element slice", t, func() {
		data := []int{55, 90}
		Convey("it should find members", func() {
			So(BinarySearch(data, 90), ShouldEqual, 1)
		})
		Convey("it should return -1 if not found", func() {
			So(BinarySearch(data, 999), ShouldEqual, -1)
		})
	})

	Convey("Given a single element slice", t, func() {
		data := []int{90}
		Convey("it should find members", func() {
			So(BinarySearch(data, 90), ShouldEqual, 0)
		})
		Convey("it should return -1 if not found", func() {
			So(BinarySearch(data, 999), ShouldEqual, -1)
		})
	})

	Convey("Given an empty slice", t, func() {
		data := []int{}
		Convey("it should return -1", func() {
			So(BinarySearch(data, 90), ShouldEqual, -1)
		})
	})

	Convey("Given nil", t, func() {
		data := []int{}
		Convey("it should return -1", func() {
			So(BinarySearch(data, 90), ShouldEqual, -1)
		})
	})
}
