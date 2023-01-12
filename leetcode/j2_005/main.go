package j2_005

func maxProduct(words []string) int {
	m := make([][]int, 0)

	res := 0
	for _, word := range words {
		bits := countBits(word)
		if bits > 0 {
			for _, exist := range m {
				if bits&exist[0] == 0 {
					res = max(res, len(word)*exist[1])
				}
			}
			m = append(m, []int{bits, len(word)})
		}
	}
	return res
}

func countBits(word string) int {
	n := 0
	for _, ch := range word {
		s := ch - 'a'
		n |= 1 << s
	}
	return n
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
