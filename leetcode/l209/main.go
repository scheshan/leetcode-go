package l209

func minSubArrayLen(target int, nums []int) int {
	l := 0
	r := 0

	cur := 0
	res := len(nums) + 1
	for r < len(nums) {
		cur += nums[r]
		r++

		for cur >= target {
			res = min(res, r-l)

			cur -= nums[l]
			l++
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
