package hash

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHash(t *testing.T) {
	Convey("given an empty hash", t, func() {
		hash := NewHash()
		Convey("it adds elements", func() {
			hash.Set("key", "value")
		})
		Convey("it remembers elements", func() {
			hash.Set("key", "value")
			So(hash.Get("key"), ShouldEqual, "value")
		})
		Convey("it remembers multiple values", func() {
			hash.Set("key", "value")
			hash.Set("key2", "value2")
			So(hash.Get("key"), ShouldEqual, "value")
			So(hash.Get("key2"), ShouldEqual, "value2")
		})
	})
}

func TestHashFunction(t *testing.T) {
	Convey("given a slice of bites", t, func() {
		data := []byte("hello hash")
		Convey("it hashes into an integer", func() {
			So(HashFunc(data), ShouldHaveSameTypeAs, uint64(1))
		})
	})
	Convey("given a sample of hash results", t, func() {
		sample_size := 1000
		sample := make(map[uint64]bool)
		for i := 0; i < sample_size; i++ {
			sample[HashFunc([]byte(strconv.Itoa(i)))] = true
		}
		Convey("it hashes into different values for different inputs", func() {
			So(len(sample), ShouldEqual, sample_size)
		})
		Convey("it produces a uniform distribution", func() {
			// TODO: look into chi square test
		})
	})
}

func BenchmarkHashFunction(b *testing.B) {
	data := []byte("hello hash")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashFunc(data)
	}
}
