package l70

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}

	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}
