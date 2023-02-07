package p1863

func subsetXORSum(nums []int) int {
	return backtrack(nums, 0, 0)
}

func backtrack(nums []int, ind int, cur int) int {
	res := cur
	for i := ind; i < len(nums); i++ {
		res += backtrack(nums, i+1, cur^nums[i])
	}
	return res
}
