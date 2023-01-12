package l2287

func rearrangeCharacters(s string, target string) int {
	tm := make(map[rune]int)
	sm := make(map[rune]int)

	for _, ch := range target {
		tm[ch]++
	}

	for _, ch := range s {
		if tm[ch] > 0 {
			sm[ch]++
		}
	}

	res := len(s)
	for ch, cnt := range tm {
		n := sm[ch] / cnt
		if n < res {
			res = n
		}
	}

	return res
}
