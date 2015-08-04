package sort

// Θ(n*log(n))
func MergeSort(data []int) []int {
	length := len(data)
	if length < 2 {
		return data
	}
	a := MergeSort(data[:length/2])
	b := MergeSort(data[length/2:])
	return merge(a, b)
}

// Θ(n)
func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	for {
		if len(a) == 0 {
			result = append(result, b...)
			break
		}
		if len(b) == 0 {
			result = append(result, a...)
			break
		}
		if a[0] <= b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}
	return result
}
