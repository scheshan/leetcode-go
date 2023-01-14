package j_21

func exchange(nums []int) []int {
	l := 0
	r := len(nums) - 1

	for l < r {
		for l < r && nums[l]&1 == 1 {
			l++
		}
		for l < r && nums[r]&1 == 0 {
			r--
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	return nums
}
