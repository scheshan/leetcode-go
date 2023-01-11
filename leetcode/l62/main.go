package l62

func uniquePaths(m int, n int) int {
	state := make([]int, n)
	for i := 0; i < n; i++ {
		state[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			state[j] = state[j] + state[j-1]
		}
	}

	return state[n-1]
}
