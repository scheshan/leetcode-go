package p1281

func subtractProductAndSum(n int) int {
	num1 := 1
	num2 := 0

	for n != 0 {
		num1 *= n % 10
		num2 += n % 10
		n /= 10
	}

	return num1 - num2
}
