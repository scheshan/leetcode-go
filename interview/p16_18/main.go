package p16_18

func patternMatching(pattern string, value string) bool {
	res := false
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	var backtrack func(int, int)
	backtrack = func(i1 int, i2 int) {
		if i1 == len(pattern) {
			if i1 == len(pattern) && i2 == len(value) {
				res = true
			}
			return
		}

		ch := pattern[i1]
		if str, ok := m1[ch]; ok {
			if i2+len(str) > len(value) {
				return
			}
			match := true
			for j := 0; j < len(str); j++ {
				if value[i2+j] != str[j] {
					match = false
					break
				}
			}
			if match {
				backtrack(i1+1, i2+len(str))
			}
		} else {
			if _, ok = m2[""]; !ok {
				m1[ch] = ""
				m2[""] = ch
				backtrack(i1+1, i2)
				delete(m1, ch)
				delete(m2, str)
			}

			for j := i2; j < len(value) && !res; j++ {
				str = value[i2 : j+1]
				if _, ok = m2[str]; ok {
					continue
				}

				m1[ch] = str
				m2[str] = ch
				backtrack(i1+1, j+1)
				delete(m1, ch)
				delete(m2, str)
			}
		}
	}
	backtrack(0, 0)
	return res
}
