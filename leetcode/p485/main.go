package p485

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	one := 0

	for _, num := range nums {
		if num == 1 {
			one++
			res = max(res, one)
		} else {
			one = 0
		}
	}

	return res
}
