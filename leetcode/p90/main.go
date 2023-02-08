package p90

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))

	sort.Sort(sort.IntSlice(nums))

	var backtrack func(ind int)
	backtrack = func(ind int) {
		arr := make([]int, len(path))
		copy(arr, path)
		res = append(res, arr)

		for i := ind; i < len(nums); i++ {
			if i > ind && nums[i-1] == nums[i] {
				continue
			}

			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
