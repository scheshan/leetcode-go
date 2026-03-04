package p41

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > len(nums) {
			nums[i] = len(nums) + 1
		}
	}

	for _, n := range nums {
		if n < 0 {
			n = -n
		}
		if n > len(nums) || nums[n-1] < 0 {
			continue
		}
		nums[n-1] *= -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
