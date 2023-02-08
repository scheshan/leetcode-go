package p557

func reverseWords(s string) string {
	res := make([]byte, len(s))

	ind := 0
	pre := 0
	r := 0
	for r < len(s) {
		for r < len(s) && s[r] != ' ' {
			r++
		}

		for i := r - 1; i >= pre; i-- {
			res[ind] = s[i]
			ind++
		}

		if r < len(s) {
			r++
			res[ind] = ' '
			ind++
		}
		pre = r
	}
	return string(res)
}
