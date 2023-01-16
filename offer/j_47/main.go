package j_47

func maxValue(grid [][]int) int {
	res := grid[0][0]
	for col := 1; col < len(grid[0]); col++ {
		grid[0][col] += grid[0][col-1]
		res = grid[0][col]
	}
	for row := 1; row < len(grid); row++ {
		grid[row][0] += grid[row-1][0]
		res = grid[row][0]
	}

	for row := 1; row < len(grid); row++ {
		for col := 1; col < len(grid[row]); col++ {
			n := grid[row-1][col]
			if grid[row][col-1] > n {
				n = grid[row][col-1]
			}
			grid[row][col] += n
			res = grid[row][col]
		}
	}

	return res
}
