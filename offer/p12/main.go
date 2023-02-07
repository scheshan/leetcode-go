package p12

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == word[0] && backtrack(board, row, col, 0, word) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, row int, col int, ind int, word string) bool {
	if board[row][col] != word[ind] {
		return false
	}

	if ind == len(word)-1 {
		return true
	}

	ch := board[row][col]
	board[row][col] = ' '
	res := false
	if row > 0 {
		res = res || backtrack(board, row-1, col, ind+1, word)
	}
	if row < len(board)-1 {
		res = res || backtrack(board, row+1, col, ind+1, word)
	}
	if col > 0 {
		res = res || backtrack(board, row, col-1, ind+1, word)
	}
	if col < len(board[row])-1 {
		res = res || backtrack(board, row, col+1, ind+1, word)
	}
	board[row][col] = ch
	return res
}
