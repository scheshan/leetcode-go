package p526

func countArrangement(n int) int {
	num := 0
	res := 0

	var backtrack func(ind int)
	backtrack = func(ind int) {
		if ind == n+1 {
			res++
			return
		}

		for i := 1; i <= n; i++ {
			if (ind%i == 0 || i%ind == 0) && num&(1<<i) == 0 {
				num ^= 1 << i
				backtrack(ind + 1)
				num ^= 1 << i
			}
		}
	}
	backtrack(1)
	return res
}
