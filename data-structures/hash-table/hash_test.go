package hash

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHashFunction(t *testing.T) {
	Convey("given a slice of bites", t, func() {
		data := []byte("hello hash")
		Convey("it hashes into an integer", func() {
			So(Hash(data), ShouldHaveSameTypeAs, 1)
		})
	})
}

func BenchmarkHashFunction(b *testing.B) {
	data := []byte("hello hash")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash(data)
	}
}
