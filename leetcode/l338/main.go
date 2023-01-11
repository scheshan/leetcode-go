package l338

func countBits(n int) []int {
	res := make([]int, n+1)
	if n > 0 {
		res[1] = 1
	}

	for i := 2; i <= n; i++ {
		pre := i >> 1
		cnt := res[pre]

		if i&1 == 1 {
			cnt++
		}
		res[i] = cnt
	}

	return res
}
