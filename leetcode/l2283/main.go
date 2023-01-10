package l2283

func digitCount(num string) bool {
	m := make([]int, 10)

	for _, ch := range num {
		m[ch-'0']++
	}

	for ind, ch := range num {
		if m[ind] != int(ch-'0') {
			return false
		}
	}
	return true
}
