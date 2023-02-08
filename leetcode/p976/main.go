package p976

import "sort"

func largestPerimeter(nums []int) int {
	sort.Sort(sort.IntSlice(nums))

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] && nums[i-1] < nums[i]+nums[i-2] && nums[i-2] < nums[i]+nums[i-1] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
