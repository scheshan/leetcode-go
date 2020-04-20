package l485

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	cur := 0

	for _, n := range nums {
		if n == 1 {
			cur++
			if cur > res {
				res = cur
			}
		} else {
			cur = 0
		}
	}

	return res
}
