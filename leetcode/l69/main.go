package l69

func mySqrt(x int) int {
	l := 1
	r := x

	res := 0
	for l <= r {
		mid := l + (r-l)>>1
		n := mid * mid

		if n == x {
			return mid
		} else if n < x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
