package hash

import (
	"crypto/sha1"
	"encoding/binary"
)

// Design:
// - Table size should be power of two
// - Use only builtin arrays (maybe slices)
// - Use bit masking to mod into table size
// - Table should dynamically resize to next power when collisions threshold
//   is reached
// - Keys will need to be stored in buckets with values for verification
//   checking and to allow for rehashing
// - Will not use open addressing

const InitialCapacity = 8

type BucketValue struct {
	key   string
	value interface{}
}

// TODO: make buckets contain multiple bucket values for collisions
type Bucket BucketValue

type Hash struct {
	buckets []Bucket
}

func NewHash() *Hash {
	return &Hash{
		buckets: make([]Bucket, InitialCapacity),
	}
}

func (h *Hash) Set(k string, v interface{}) {
	index := truncate(HashFunc([]byte(k)), len(h.buckets))
	h.buckets[index] = Bucket{
		key:   k,
		value: v,
	}
}

func (h *Hash) Get(k string) interface{} {
	index := truncate(HashFunc([]byte(k)), len(h.buckets))
	// TODO: verify key to detect collisions
	return h.buckets[index].value
}

func HashFunc(data []byte) uint64 {
	hash := sha1.New()
	hash.Write(data)
	sum := make([]byte, 0)
	sum = hash.Sum(sum)

	result := binary.BigEndian.Uint64(sum)
	return result
}

func truncate(value uint64, size int) uint64 {
	return value % uint64(size)
}

// TODO: use masking vs mod

// 000000000000 1111
// 101010101101 0101  bitwise and
// 000000000000 0101
