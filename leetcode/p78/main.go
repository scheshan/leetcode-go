package p78

func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	backtrack(nums, 0, &path, &res)
	return res
}

func backtrack(nums []int, ind int, path *[]int, res *[][]int) {
	arr := make([]int, len(*path))
	copy(arr, *path)
	*res = append(*res, arr)

	if ind == len(nums) {
		return
	}

	for i := ind; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backtrack(nums, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
