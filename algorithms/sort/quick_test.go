package sort

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestQuickSort(t *testing.T) {
	Convey("given an unordered slice of ints", t, func() {
		data := []int{7, 82, 3, 5, 1, 111}
		Convey("it properly sorts the list", func() {
			data := QuickSort(data)
			So(data, ShouldResemble, []int{1, 3, 5, 7, 82, 111})
		})
	})
	Convey("given an empty slice of ints", t, func() {
		data := []int{}
		Convey("it does nothing", func() {
			data := QuickSort(data)
			So(data, ShouldBeEmpty)
		})
	})
	Convey("given a single element slice of ints", t, func() {
		data := []int{55}
		Convey("it does nothing", func() {
			data := QuickSort(data)
			So(data, ShouldResemble, []int{55})
		})
	})
}
