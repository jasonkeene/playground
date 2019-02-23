package hash_test

import (
	"strconv"
	"testing"

	hash "github.com/jasonkeene/playground/data-structures/hash-table"
)

func TestHash(t *testing.T) {
	h := hash.NewHash()

	h.Set("key", "value")
	h.Set("key2", "value2")

	if h.Get("key") != "value" {
		t.Fatal("key != value")
	}
	if h.Get("key2") != "value2" {
		t.Fatal("key2 != value2")
	}
}

func TestHashFunction(t *testing.T) {
	const sample_size = 1000
	sample := make(map[uint64]bool, sample_size)
	for i := 0; i < sample_size; i++ {
		sample[hash.HashFunc([]byte(strconv.Itoa(i)))] = true
	}

	if len(sample) != sample_size {
		t.Fatalf("hash func did not have uniqe values for each sample")
	}
	// TODO: look into chi square test for uniform distribution
}

func BenchmarkHashFunction(b *testing.B) {
	data := []byte("hello hash")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash.HashFunc(data)
	}
}
