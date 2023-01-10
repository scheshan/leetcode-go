package l509

func fib(n int) int {
	if n < 2 {
		return n
	}

	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}
