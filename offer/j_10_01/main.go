package j_10_01

func fib(n int) int {
	if n < 2 {
		return n
	}

	n1 := 0
	n2 := 1

	for i := 2; i <= n; i++ {
		n1, n2 = n2, (n1+n2)%1000000007
	}
	return n2
}
