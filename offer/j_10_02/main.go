package j_10_02

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}

	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		n1, n2 = n2, (n1+n2)%1000000007
	}

	return n2
}
