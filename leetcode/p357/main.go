package p357

func countNumbersWithUniqueDigits(n int) int {
	visit := 0
	res := 0

	var backtrack func(ind int, num int)
	backtrack = func(ind int, num int) {
		res++
		if ind == n {
			return
		}

		for i := 0; i < 10; i++ {
			if ind == 0 && i == 0 {
				continue
			}

			if visit&(1<<i) != 0 {
				continue
			}
			visit |= 1 << i
			backtrack(ind+1, num*10+i)
			visit ^= 1 << i
		}
	}
	backtrack(0, 0)
	return res
}
