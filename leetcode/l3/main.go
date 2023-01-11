package l3

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)

	l := 0
	r := 0

	res := 0

	for r < len(s) {
		m[s[r]] = m[s[r]] + 1
		for len(m) != r-l+1 {
			m[s[l]] = m[s[l]] - 1
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}

		r++
	}

	return res
}
