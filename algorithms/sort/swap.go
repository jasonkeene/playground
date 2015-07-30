package sort

// Θ(1)
func Swap(data []int, i, j int) {
	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}
