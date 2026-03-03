package p189

func rotate(nums []int, k int) {
	k = k % len(nums)
	rotate0(nums, 0, len(nums)-1)
	rotate0(nums, 0, k-1)
	rotate0(nums, k, len(nums)-1)
}

func rotate0(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
