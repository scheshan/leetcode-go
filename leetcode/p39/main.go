package p39

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	var backtrack func(nums []int, ind int, total int)
	backtrack = func(nums []int, ind int, total int) {
		if total <= 0 {
			if total == 0 {
				arr := make([]int, len(path))
				copy(arr, path)
				res = append(res, arr)
			}
			return
		}

		for i := ind; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(nums, i, total-nums[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(candidates, 0, target)
	return res
}
