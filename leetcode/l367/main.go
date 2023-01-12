package l367

func isPerfectSquare(num int) bool {
	l := 1
	r := num + 1

	for l < r {
		mid := l + (r-l)>>1
		n := mid * mid
		if n == num {
			return true
		} else if n < num {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}
