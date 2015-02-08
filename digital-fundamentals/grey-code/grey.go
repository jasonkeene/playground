package grey

// Convert from binary representation to grey code. O(1)
func ToGrey(x uint) uint {
	return x>>1 ^ x
}

// Convert from grey code to binary representation. O(n) where n is the place
// of the most significant binary digit.
func FromGrey(x uint) uint {
	for mask := x >> 1; mask != 0; mask = mask >> 1 {
		x = x ^ mask
	}
	return x
}
