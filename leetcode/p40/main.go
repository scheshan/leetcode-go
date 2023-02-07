package p40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))

	path := make([]int, 0)
	res := make([][]int, 0)

	var backtrack func([]int, int, int)
	backtrack = func(nums []int, ind int, total int) {
		if total <= 0 {
			if total == 0 {
				arr := make([]int, len(path))
				copy(arr, path)
				res = append(res, arr)
			}
			return
		}

		for i := ind; i < len(nums) && nums[i] <= total; i++ {
			if i > ind && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(nums, i+1, total-nums[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(candidates, 0, target)
	return res
}
