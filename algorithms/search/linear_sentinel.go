package search

// Θ(n)
func LinearSentinelSearch(data []int, number int) int {
	lastIndex := len(data) - 1
	last := data[lastIndex]
	data[lastIndex] = number
	var i int
	for data[i] != number {
		i++
	}
	data[lastIndex] = last
	if i < lastIndex || data[i] == number {
		return i
	}
	return -1
}
