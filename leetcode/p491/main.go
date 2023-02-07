package p491

func findSubsequences(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)

	var backtrack func(ind int)
	backtrack = func(ind int) {
		if len(path) >= 2 {
			arr := make([]int, len(path))
			copy(arr, path)
			res = append(res, arr)
		}

		for i := ind; i < len(nums); i++ {
			duplicate := false
			for j := ind; j < i; j++ {
				if nums[j] == nums[i] {
					duplicate = true
					break
				}
			}
			if duplicate {
				continue
			}

			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return res
}
