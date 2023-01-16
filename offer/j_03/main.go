package j_03

func findRepeatNumber(nums []int) int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num == len(nums) {
			num = 0
		}

		if nums[num] < 0 {
			return num
		}

		if nums[num] == 0 {
			nums[num] = -len(nums)
		} else {
			nums[num] = -nums[num]
		}
	}
	return -1
}
