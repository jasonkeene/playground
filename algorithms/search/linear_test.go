package search

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestLinearSearchCorrectness(t *testing.T) {
	Convey("Given an unordered slice of integers", t, func() {
		data := []int{0, 1, 6, 90, 3, 42, 71, 9, 53, 567}
		Convey("it should find members", func() {
			So(LinearSearch(data, 90), ShouldEqual, 3)
		})
		Convey("it should return -1 if not found", func() {
			So(LinearSearch(data, 999), ShouldEqual, -1)
		})
	})
}

func BenchmarkLinearBestCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LinearSearch(benchmarkData, 8191711)
	}
}

func BenchmarkLinearAverageCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LinearSearch(benchmarkData, 8348554)
	}
}

func BenchmarkLinearWorstCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LinearSearch(benchmarkData, 6309519)
	}
}
