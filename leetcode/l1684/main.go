package l1684

func countConsistentStrings(allowed string, words []string) int {
	m := make([]bool, 26)

	for _, c := range allowed {
		m[c-'a'] = true
	}

	res := 0
	for _, word := range words {
		match := true
		for _, c := range word {
			if !m[c-'a'] {
				match = false
				break
			}
		}

		if match {
			res++
		}
	}
	return res
}
