package p37

func solveSudoku(board [][]byte) {
	rows := make([]int, 9)
	cols := make([]int, 9)
	blocks := make([]int, 9)
	blanks := make([][]int, 0)
	complete := false

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] >= '1' && board[row][col] <= '9' {
				num := board[row][col] - '0'
				rows[row] |= 1 << num
				cols[col] |= 1 << num
				blocks[getBlockInd(row, col)] = blocks[getBlockInd(row, col)] | (1 << num)
			} else {
				blanks = append(blanks, []int{row, col})
			}
		}
	}

	var backtrack func(ind int)
	backtrack = func(ind int) {
		if ind == len(blanks) {
			complete = true
			return
		}

		row, col := blanks[ind][0], blanks[ind][1]
		blockInd := getBlockInd(row, col)

		for i := 1; i <= 9 && !complete; i++ {
			n := 1 << i
			if rows[row]&n != 0 || cols[col]&n != 0 || blocks[blockInd]&n != 0 {
				continue
			}
			rows[row] |= n
			cols[col] |= n
			blocks[blockInd] |= n
			board[row][col] = '0' + byte(i)
			backtrack(ind + 1)

			if !complete {
				rows[row] ^= n
				cols[col] ^= n
				blocks[blockInd] ^= n
				board[row][col] = '.'
			}
		}
	}
	backtrack(0)
}

func getBlockInd(row int, col int) int {
	return (row / 3 * 3) + (col / 3)
}
