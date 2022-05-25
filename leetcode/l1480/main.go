package l1480

func runningSum(nums []int) []int {
	pre := 0
	for i, n := range nums {
		pre += n
		nums[i] = pre
	}
	return nums
}
