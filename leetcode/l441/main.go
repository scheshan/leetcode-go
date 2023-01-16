package main

func arrangeCoins(n int) int {
	l := 1
	r := n

	res := 0

	for l <= r {
		mid := l + (r-l)>>1

		cnt := (1 + mid) * mid >> 1
		if cnt == n {
			res = mid
			break
		} else if cnt > n {
			res = mid - 1
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {
	arrangeCoins(5)
}
