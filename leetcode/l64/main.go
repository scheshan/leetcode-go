package l64

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[len(grid)-1][len(grid[len(grid)-1])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
