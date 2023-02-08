package p202

func isHappy(n int) bool {
	visit := make(map[int]bool)

	for n != 0 {
		sum := 0
		n1 := n
		for n1 != 0 {
			sum += (n1 % 10) * (n1 % 10)
			n1 /= 10
		}

		if sum == 1 {
			return true
		}

		if visit[sum] {
			return false
		}
		visit[sum] = true
		n = sum
	}

	return false
}
