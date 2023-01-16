package j_57

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	for l < r {
		n := nums[l] + nums[r]
		if n == target {
			return []int{nums[l], nums[r]}
		} else if n < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
