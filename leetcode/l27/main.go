package l27

func removeElement(nums []int, val int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		if nums[l] != val {
			l++
		} else {
			nums[l] = nums[r]
			r--
		}
	}

	return l
}
