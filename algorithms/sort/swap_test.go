package sort

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestSwap(t *testing.T) {
	Convey("given an unordered slice of ints", t, func() {
		data := []int{7, 3, 5, 1}
		Convey("it swaps", func() {
			Swap(data, 1, 3)
			So(data, ShouldResemble, []int{7, 1, 5, 3})
		})
	})
}
