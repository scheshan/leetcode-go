package l26

func removeDuplicates(nums []int) int {
	slow := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			slow++
			nums[slow] = nums[i]
		}
	}

	return slow + 1
}
