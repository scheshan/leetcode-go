package p290

func wordPattern(pattern string, s string) bool {
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	l := 0

	pre := 0
	r := 0

	for r < len(s) && l < len(pattern) {
		for r < len(s) && s[r] != ' ' {
			r++
		}
		str := s[pre:r]

		s2, o1 := m1[pattern[l]]
		b, o2 := m2[str]

		if !o1 && !o2 {
			m1[pattern[l]] = str
			m2[str] = pattern[l]
		} else if o1 && o2 {
			if s2 != str || b != pattern[l] {
				return false
			}
		} else {
			return false
		}

		if r < len(s) {
			r++
		}
		pre = r
		l++
	}

	return l == len(pattern) && r == len(s)
}
