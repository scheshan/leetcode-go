package l2351

func repeatedCharacter(s string) byte {
	m := make([]bool, 26)

	for _, r := range s {
		if m[r-'a'] {
			return byte(r)
		}
		m[r-'a'] = true
	}
	return byte(' ')
}
