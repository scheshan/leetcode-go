package p140

import "strings"

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}

	path := make([]string, 0)
	res := make([]string, 0)

	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == len(s) {
			b := strings.Builder{}
			for _, p := range path {
				if b.Len() > 0 {
					b.WriteByte(' ')
				}
				b.WriteString(p)
			}
			res = append(res, b.String())
			return
		}

		for i := ind; i < len(s); i++ {
			p := s[ind : i+1]
			if m[p] {
				path = append(path, p)
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)

	return res
}
