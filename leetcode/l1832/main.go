package l1832

func checkIfPangram(sentence string) bool {
	cnt := 0
	m := make([]bool, 26)

	for _, ch := range sentence {
		if !m[ch-'a'] {
			m[ch-'a'] = true
			cnt++
		}
	}
	return cnt == 26
}
