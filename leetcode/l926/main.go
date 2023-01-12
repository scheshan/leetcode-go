package l926

func minFlipsMonoIncr(s string) int {
	n0 := 0
	n1 := 0

	for i, ch := range s {
		if i == 0 {
			if ch == '0' {
				n1 = 1
			} else {
				n0 = 1
			}
		} else {
			if ch == '0' {
				n1 = min(n0, n1) + 1
			} else {
				n0, n1 = n0+1, min(n0, n1)
			}
		}
	}

	return min(n0, n1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
