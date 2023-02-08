package p1219

func getMaximumGold(grid [][]int) int {
	res := 0

	var backtrack func(int, int, int)
	backtrack = func(row int, col int, cur int) {
		if grid[row][col] == 0 {
			return
		}

		num := grid[row][col]
		cur += num
		if cur > res {
			res = cur
		}
		grid[row][col] = 0

		if row > 0 {
			backtrack(row-1, col, cur)
		}
		if row < len(grid)-1 {
			backtrack(row+1, col, cur)
		}
		if col > 0 {
			backtrack(row, col-1, cur)
		}
		if col < len(grid[row])-1 {
			backtrack(row, col+1, cur)
		}
		grid[row][col] = num
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] > 0 {
				backtrack(row, col, 0)
			}
		}
	}
	return res
}
