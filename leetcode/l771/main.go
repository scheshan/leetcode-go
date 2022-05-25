package l771

func numJewelsInStones(J string, S string) int {
	m := make([]bool, 80)

	for _, ch := range J {
		m[ch-'a'] = true
	}

	res := 0
	for _, ch := range S {
		if m[ch-'a'] {
			res++
		}
	}
	return res
}
