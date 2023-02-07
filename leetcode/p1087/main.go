package p1087

import "sort"

func expand(s string) []string {
	path := make([]byte, 0, len(s))
	res := make([]string, 0)

	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == len(s) {
			res = append(res, string(path))
			return
		}

		start := ind
		end := ind + 1

		if s[ind] == '{' {
			start = ind + 1
			for end < len(s) && s[end] != '}' {
				end++
			}
			end++
		}

		for i := start; i < end; i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				path = append(path, s[i])
				backtrack(end)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)

	sort.Sort(sort.StringSlice(res))

	return res
}
