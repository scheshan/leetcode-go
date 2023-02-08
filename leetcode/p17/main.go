package p17

func letterCombinations(digits string) []string {
	chs := buildMap()

	res := make([]string, 0)
	path := make([]byte, 0, len(digits))
	var backtrack func(ind int)
	backtrack = func(ind int) {
		if ind == len(digits) {
			if ind > 0 {
				res = append(res, string(path))
			}
			return
		}

		arr := chs[digits[ind]-'0']
		for _, b := range arr {
			path = append(path, b)
			backtrack(ind + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

func buildMap() [][]byte {
	res := [][]byte{
		{},
		{},
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	return res
}
