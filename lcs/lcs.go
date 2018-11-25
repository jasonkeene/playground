package lcs

func Recursive(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}

	if a[0] == b[0] {
		result := a[0:1]
		if len(a) == 1 || len(b) == 1 {
			return result
		}
		return result + Recursive(a[1:], b[1:])
	}

	var resultA, resultB string
	if len(a) != 1 {
		resultA = Recursive(a[1:], b)
	}
	if len(b) != 1 {
		resultB = Recursive(a, b[1:])
	}
	if len(resultA) > len(resultB) {
		return resultA
	}
	return resultB
}

type MemKey struct {
	A string
	B string
}

func RecursiveMemoized(a, b string, mem map[MemKey]string) string {
	if r, ok := mem[MemKey{a, b}]; ok {
		return r
	}

	if len(a) == 0 || len(b) == 0 {
		mem[MemKey{a, b}] = ""
		return ""
	}

	if a[0] == b[0] {
		result := a[0:1]
		if len(a) == 1 || len(b) == 1 {
			mem[MemKey{a, b}] = result
			return result
		}
		result += RecursiveMemoized(a[1:], b[1:], mem)
		mem[MemKey{a, b}] = result
		return result
	}

	var resultA, resultB string
	if len(a) != 1 {
		resultA = RecursiveMemoized(a[1:], b, mem)
	}
	if len(b) != 1 {
		resultB = RecursiveMemoized(a, b[1:], mem)
	}
	if len(resultA) > len(resultB) {
		mem[MemKey{a, b}] = resultA
		return resultA
	}
	mem[MemKey{a, b}] = resultB
	return resultB
}

func RecursiveIndexes(a, b string, i, j int) string {
	if len(a) == i || len(b) == j {
		return ""
	}

	if a[i] == b[j] {
		return a[i:i+1] + RecursiveIndexes(a, b, i+1, j+1)
	}

	var resultA, resultB string
	if len(a) != i+1 {
		resultA = RecursiveIndexes(a, b, i+1, j)
	}
	if len(b) != j+1 {
		resultB = RecursiveIndexes(a, b, i, j+1)
	}
	if len(resultA) > len(resultB) {
		return resultA
	}
	return resultB
}

type MemKeyIndexes struct {
	I int
	J int
}

func RecursiveIndexesMemoized(a, b string, i, j int, mem map[MemKeyIndexes]string) string {
	if r, ok := mem[MemKeyIndexes{i, j}]; ok {
		return r
	}

	if len(a) == i || len(b) == j {
		mem[MemKeyIndexes{i, j}] = ""
		return ""
	}

	if a[i] == b[j] {
		result := a[i:i+1] + RecursiveIndexesMemoized(a, b, i+1, j+1, mem)
		mem[MemKeyIndexes{i, j}] = result
		return result
	}

	var resultA, resultB string
	if len(a) != i+1 {
		resultA = RecursiveIndexesMemoized(a, b, i+1, j, mem)
	}
	if len(b) != j+1 {
		resultB = RecursiveIndexesMemoized(a, b, i, j+1, mem)
	}
	if len(resultA) > len(resultB) {
		mem[MemKeyIndexes{i, j}] = resultA
		return resultA
	}
	mem[MemKeyIndexes{i, j}] = resultB
	return resultB
}

type DiffResult struct {
	A string
	B string
}

func Diff(a, b string) DiffResult {
	lcs := RecursiveIndexesMemoized(a, b, 0, 0, map[MemKeyIndexes]string{})
	return DiffResult{
		A: genP(a, lcs, "-"),
		B: genP(b, lcs, "+"),
	}
}

func genP(p, lcs, op string) string {
	var (
		result string
		s      int
		d      bool
	)
	for i := 0; i < len(p); i++ {
		if p[i] == lcs[s] {
			if d {
				d = false
				result += ">"
			}
			result += p[i : i+1]
			s++
		} else {
			if !d {
				d = true
				result += op + "<"
			}
			result += p[i : i+1]
		}
	}
	return result
}
