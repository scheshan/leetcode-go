package l53

func maxSubArray(nums []int) int {
	min := 0
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]

		if nums[i] > res {
			res = nums[i]
		}
		if nums[i]-nums[min] > res {
			res = nums[i] - nums[min]
		}
		if nums[i] < nums[min] {
			min = i
		}
	}
	return res
}
