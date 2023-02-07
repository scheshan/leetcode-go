package p086

func partition(s string) [][]string {
	path := make([]string, 0, len(s))
	res := make([][]string, 0)

	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == len(s) {
			arr := make([]string, len(path))
			copy(arr, path)
			res = append(res, arr)
			return
		}

		for i := ind; i < len(s); i++ {
			if valid(s, ind, i) {
				path = append(path, s[ind:i+1])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func valid(str string, begin int, end int) bool {
	for begin < end {
		if str[begin] != str[end] {
			return false
		}
		begin++
		end--
	}
	return true
}
