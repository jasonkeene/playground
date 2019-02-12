package str

import "fmt"

var cache = make(map[string]string)

func LCS(a, b string) string {
	result, ok := cache[key(a, b)]
	if ok {
		return result
	}

	fmt.Printf("solving: %q %q\n", a, b)
	if len(a) == 0 || len(b) == 0 {
		cache[key(a, b)] = ""
		return ""
	}

	if a[len(a)-1] == b[len(b)-1] {
		result = LCS(a[:len(a)-1], b[:len(b)-1]) + string(b[len(b)-1])
		cache[key(a, b)] = result
		return result
	}
	alcs := LCS(a[:len(a)-1], b)
	blcs := LCS(a, b[:len(b)-1])
	if alcs > blcs {
		cache[key(a, b)] = alcs
		return alcs
	}
	cache[key(a, b)] = blcs
	return blcs
}

func key(a, b string) string {
	return a + "|" + b
}
