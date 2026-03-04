package p1732

func largestAltitude(gain []int) int {
	res := 0
	sum := 0

	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		res = max(res, sum)
	}

	return res
}
