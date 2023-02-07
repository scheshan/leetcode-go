package p47

import "sort"

func permuteUnique(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)
	visit := make([]bool, len(nums))

	sort.Sort(sort.IntSlice(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			arr := make([]int, len(path))
			copy(arr, path)
			res = append(res, arr)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visit[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !visit[i-1] {
				continue
			}
			path = append(path, nums[i])
			visit[i] = true
			backtrack()
			visit[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return res
}
