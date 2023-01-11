package l724

func pivotIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	total := nums[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		left := 0
		if i > 0 {
			left = nums[i-1]
		}
		if total-nums[i] == left {
			return i
		}
	}
	return -1
}
