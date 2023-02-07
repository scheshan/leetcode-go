package p52

func totalNQueens(n int) int {
	grid := make([][]bool, n)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, n)
	}
	cols := make([]bool, n)
	res := 0

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || !valid(grid, row, col) {
				continue
			}

			cols[col] = true
			grid[row][col] = true
			backtrack(row + 1)
			grid[row][col] = false
			cols[col] = false
		}
	}
	backtrack(0)
	return res
}

func valid(grid [][]bool, row int, col int) bool {
	r1 := row
	c1 := col

	for c1 >= 0 && r1 >= 0 {
		if grid[r1][c1] {
			return false
		}
		c1--
		r1--
	}

	r2 := row
	c2 := col
	for c2 < len(grid[row]) && r2 >= 0 {
		if grid[r2][c2] {
			return false
		}
		c2++
		r2--
	}
	return true
}
