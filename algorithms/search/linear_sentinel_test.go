package search

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestLinearSentinelSearchCorrectness(t *testing.T) {
	Convey("Given an unordered slice of integers", t, func() {
		data := []int{0, 1, 6, 90, 3, 42, 71, 9, 53, 567}
		Convey("it should find members", func() {
			So(LinearSentinelSearch(data, 90), ShouldEqual, 3)
		})
		Convey("it should return -1 if not found", func() {
			So(LinearSentinelSearch(data, 999), ShouldEqual, -1)
		})
	})
}

func BenchmarkLinearSentinelBestCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LinearSentinelSearch(benchmarkData, 8191711)
	}
}

func BenchmarkLinearSentinelAverageCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LinearSentinelSearch(benchmarkData, 8348554)
	}
}

func BenchmarkLinearSentinelWorstCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LinearSentinelSearch(benchmarkData, 6309519)
	}
}
