package l1876

func countGoodSubstrings(s string) int {
	l := 0
	r := 2

	res := 0
	for r < len(s) {
		if s[l] != s[l+1] && s[l+1] != s[l+2] && s[l] != s[l+2] {
			res++
		}
		r++
		l++
	}

	return res
}
