package sort

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestSelectionSort(t *testing.T) {
	Convey("given an unordered slice of ints", t, func() {
		data := []int{7, 82, 3, 5, 1, 111}
		Convey("it properly sorts the list", func() {
			SelectionSort(data)
			So(data, ShouldResemble, []int{1, 3, 5, 7, 82, 111})
		})
	})
	Convey("given an empty slice of ints", t, func() {
		data := []int{}
		Convey("it does nothing", func() {
			SelectionSort(data)
			So(data, ShouldBeEmpty)
		})
	})
	Convey("given a single element slice of ints", t, func() {
		data := []int{55}
		Convey("it does nothing", func() {
			SelectionSort(data)
			So(data, ShouldResemble, []int{55})
		})
	})
}

func TestSwap(t *testing.T) {
	Convey("given an unordered slice of ints", t, func() {
		data := []int{7, 3, 5, 1}
		Convey("it swaps", func() {
			Swap(data, 1, 3)
			So(data, ShouldResemble, []int{7, 1, 5, 3})
		})
	})
}
