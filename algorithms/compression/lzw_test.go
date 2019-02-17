package compression_test

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/algorithms/compression"
)

func TestLZW(t *testing.T) {
	text := []byte("TATAGATCTTAATATA")
	compressed := compression.LZW(text)
	expectedCompressed := []int{
		'T',
		'A',
		256, // TA
		'G',
		257, // AT
		'C',
		'T',
		256, // TA
		257, // AT
		264, // ATA
	}

	if !cmp.Equal(compressed, expectedCompressed) {
		t.Fatal(cmp.Diff(compressed, expectedCompressed))
	}

	decompressed := compression.LZWDecompress(compressed)
	if !cmp.Equal(decompressed, text) {
		t.Fatal(cmp.Diff(decompressed, text))
	}
}

func TestLZWCorrectness(t *testing.T) {
	testCompressionCorrectness(t, compression.LZW, compression.LZWDecompress)
}

func testCompressionCorrectness(
	t *testing.T,
	comp func([]byte) []int,
	dec func([]int) []byte,
) {
	for i := 0; i < 1024; i++ {
		tc := []byte(randomString(i))
		result := dec(comp(tc))
		if !cmp.Equal(tc, result) {
			t.Fatal(cmp.Diff(tc, result))
		}
	}
}

func randomString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += string(rand.Intn(26) + 1 | 0x40)
	}
	return s
}
