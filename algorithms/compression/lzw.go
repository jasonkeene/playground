package compression

func LZW(in []byte) []int {
	if len(in) == 0 {
		return []int{}
	}

	table, next := initTable()

	var (
		s     string
		match string
		out   []int
	)
	for _, c := range in {
		s += string(c)
		if _, ok := table[s]; ok {
			continue
		}

		table[s] = next
		next++

		match, s = s[:len(s)-1], string(s[len(s)-1])
		out = append(out, table[match])
	}

	return append(out, table[s])
}

func initTable() (map[string]int, int) {
	table := make(map[string]int, 256)
	for i := 0; i < 256; i++ {
		table[string(i)] = i
	}
	return table, 256
}

func initDecompressTable() (map[int]string, int) {
	table := make(map[int]string, 256)
	for i := 0; i < 256; i++ {
		table[i] = string(i)
	}
	return table, 256
}

func LZWDecompress(in []int) []byte {
	if len(in) == 0 {
		return []byte{}
	}

	table, next := initDecompressTable()

	var (
		out  []byte
		prev string
	)

	for i, currCp := range in {
		s, ok := table[currCp]
		if !ok {
			s = prev + string(prev[0])
		}
		out = append(out, []byte(s)...)

		if i == 0 {
			prev = s
			continue
		}

		table[next] = prev + string(s[0])
		next++
		prev = s
	}

	return out
}
