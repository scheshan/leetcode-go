package p784

func letterCasePermutation(s string) []string {
	path := make([]byte, 0, len(s))
	res := make([]string, 0)

	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == len(s) {
			res = append(res, string(path))
			return
		}

		path = append(path, s[ind])
		backtrack(ind + 1)
		path = path[:len(path)-1]

		if s[ind] >= 'a' && s[ind] <= 'z' {
			path = append(path, s[ind]-32)
			backtrack(ind + 1)
			path = path[:len(path)-1]
		} else if s[ind] >= 'A' && s[ind] <= 'Z' {
			path = append(path, s[ind]+32)
			backtrack(ind + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
