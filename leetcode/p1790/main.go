package p1790

func areAlmostEqual(s1 string, s2 string) bool {
	l := -1
	r := -1

	for i := range s1 {
		if s1[i] != s2[i] {
			if l < 0 {
				l = i
			} else if r < 0 {
				r = i
			} else {
				return false
			}
		}
	}

	return l < 0 || (r >= 0 && s1[l] == s2[r] && s1[r] == s2[l])
}
