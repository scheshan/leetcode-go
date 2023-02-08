package p216

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, k)

	var backtrack func(int, int)
	backtrack = func(ind int, sum int) {
		if len(path) == k {
			if sum == n {
				arr := make([]int, len(path))
				copy(arr, path)
				res = append(res, arr)
			}
			return
		}

		for i := ind; i <= 9; i++ {
			if sum+i > n {
				break
			}

			path = append(path, i)
			backtrack(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, 0)
	return res
}
