package p980

func uniquePathsIII(grid [][]int) int {
	res := 0
	cnt := 0

	startRow := 0
	startCol := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				cnt++
			} else if grid[row][col] == 1 {
				startRow = row
				startCol = col
			}
		}
	}

	var backtrack func(row int, col int)
	backtrack = func(row int, col int) {
		if grid[row][col] == 2 {
			if cnt == 0 {
				res++
			}
			return
		}
		if grid[row][col] < 0 {
			return
		}

		num := grid[row][col]
		grid[row][col] = -1

		if num == 0 {
			cnt--
		}

		if row > 0 {
			backtrack(row-1, col)
		}
		if row < len(grid)-1 {
			backtrack(row+1, col)
		}
		if col > 0 {
			backtrack(row, col-1)
		}
		if col < len(grid[row])-1 {
			backtrack(row, col+1)
		}

		if num == 0 {
			cnt++
		}
		grid[row][col] = num
	}

	backtrack(startRow, startCol)
	return res
}
