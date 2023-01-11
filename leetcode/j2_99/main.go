package main

func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			n := grid[i-1][j]
			if j >= 1 && grid[i][j-1] < n {
				n = grid[i][j-1]
			}
			grid[i][j] += n
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	minPathSum([][]int{{1, 2}, {1, 1}})
}
