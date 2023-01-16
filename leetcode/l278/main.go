package l278

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	l := 1
	r := n

	res := 0

	for l <= r {
		mid := l + (r-l)>>1
		if isBadVersion(mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}
