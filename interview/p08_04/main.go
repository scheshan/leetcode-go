package p08_04

import "sort"

func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)

	sort.Sort(sort.IntSlice(nums))

	var backtrack func(int)
	backtrack = func(ind int) {
		arr := make([]int, len(path))
		copy(arr, path)
		res = append(res, arr)

		if ind == len(nums) {
			return
		}

		for i := ind; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
