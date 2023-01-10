package l374

func guess(n int) int {
	return 0
}

func guessNumber(n int) int {
	l := 1
	r := n
	for l <= r {
		mid := l + (r-l)>>1

		g := guess(mid)
		if g == 0 {
			return mid
		} else if g > 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
