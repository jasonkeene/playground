package str

func LCS(a, b string) string {
	cache := make(map[string]string)
	return lcs(a, b, cache)
}

func lcs(a, b string, cache map[string]string) string {
	v, ok := cache[key(a, b)]
	if ok {
		return v
	}

	if len(a) == 0 || len(b) == 0 {
		cache[key(a, b)] = ""
		return ""
	}

	if a[len(a)-1] == b[len(b)-1] {
		v = lcs(a[:len(a)-1], b[:len(b)-1], cache) + string(b[len(b)-1])
		cache[key(a, b)] = v
		return v
	}
	av := lcs(a[:len(a)-1], b, cache)
	bv := lcs(a, b[:len(b)-1], cache)
	if av > bv {
		cache[key(a, b)] = av
		return av
	}
	cache[key(a, b)] = bv
	return bv
}

func key(a, b string) string {
	return a + "|" + b
}

func LCSTable(a, b string) string {
	table := LCSLengthTable(a, b)
	return assembleLCS(a, b, table, len(a)-1, len(b)-1)
}

func assembleLCS(a, b string, table [][]int, i, j int) string {
	if value(table, i, j) == 0 {
		return ""
	}

	if a[i] == b[j] {
		return assembleLCS(a, b, table, i-1, j-1) + string(a[i])
	}

	if above(table, i, j) > left(table, i, j) {
		return assembleLCS(a, b, table, i-1, j)
	}

	return assembleLCS(a, b, table, i, j-1)
}

func AllLCS(a, b string) map[string]struct{} {
	results := make(map[string]struct{})
	table := LCSLengthTable(a, b)
	assembleAllLCS(a, b, "", table, results, len(a)-1, len(b)-1)
	return results
}

func assembleAllLCS(a, b, partial string, table [][]int, results map[string]struct{}, i, j int) {
	if value(table, i, j) == 0 {
		results[partial] = struct{}{}
		return
	}

	if a[i] == b[j] {
		assembleAllLCS(a, b, string(a[i])+partial, table, results, i-1, j-1)
		return
	}

	u := above(table, i, j)
	l := left(table, i, j)

	if u > l {
		assembleAllLCS(a, b, partial, table, results, i-1, j)
		return
	}

	if u < l {
		assembleAllLCS(a, b, partial, table, results, i, j-1)
		return
	}

	assembleAllLCS(a, b, partial, table, results, i-1, j)
	assembleAllLCS(a, b, partial, table, results, i, j-1)
}

func LCSLengthTable(a, b string) [][]int {
	table := make([][]int, len(a))
	for i := range table {
		table[i] = make([]int, len(b))

		for j := range table[i] {
			if a[i] == b[j] {
				table[i][j] = diagonal(table, i, j) + 1
				continue
			}

			table[i][j] = max(above(table, i, j), left(table, i, j))
		}
	}
	return table
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func value(table [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return table[i][j]
}

func diagonal(table [][]int, i, j int) int {
	return value(table, i-1, j-1)
}

func above(table [][]int, i, j int) int {
	return value(table, i-1, j)
}

func left(table [][]int, i, j int) int {
	return value(table, i, j-1)
}
