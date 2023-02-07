package p08_12

func solveNQueens(n int) [][]string {
	grid := make([][]byte, n)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = '.'
		}
	}
	cols := make([]bool, n)
	res := make([][]string, 0)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			arr := make([]string, n)
			for i := 0; i < n; i++ {
				arr[i] = string(grid[i])
			}
			res = append(res, arr)
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || !valid(grid, row, col) {
				continue
			}

			cols[col] = true
			grid[row][col] = 'Q'
			backtrack(row + 1)
			grid[row][col] = '.'
			cols[col] = false
		}
	}
	backtrack(0)
	return res
}

func valid(grid [][]byte, row int, col int) bool {
	r1 := row
	c1 := col

	for c1 >= 0 && r1 >= 0 {
		if grid[r1][c1] == 'Q' {
			return false
		}
		c1--
		r1--
	}

	r2 := row
	c2 := col
	for c2 < len(grid[row]) && r2 >= 0 {
		if grid[r2][c2] == 'Q' {
			return false
		}
		c2++
		r2--
	}
	return true
}
