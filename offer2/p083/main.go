package p083

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	backtrack(nums, 0, &res)
	return res
}

func backtrack(nums []int, ind int, res *[][]int) {
	if ind == len(nums) {
		arr := make([]int, len(nums))
		copy(arr, nums)
		*res = append(*res, arr)
		return
	}

	for i := ind; i < len(nums); i++ {
		nums[ind], nums[i] = nums[i], nums[ind]
		backtrack(nums, ind+1, res)
		nums[ind], nums[i] = nums[i], nums[ind]
	}
}
