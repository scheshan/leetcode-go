package l130

import "container/list"

func addQueue(queue *list.List, board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] != 'O' {
		return
	}

	board[i][j] = 'Y'

	queue.PushBack(i)
	queue.PushBack(j)
}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	queue := list.New()

	for j := 0; j < len(board[0]); j++ {
		addQueue(queue, board, 0, j)
		addQueue(queue, board, len(board)-1, j)
	}
	for i := 1; i < len(board)-1; i++ {
		addQueue(queue, board, i, 0)
		addQueue(queue, board, i, len(board[i])-1)
	}

	var i, j int
	for queue.Len() > 0 {
		i = queue.Front().Value.(int)
		queue.Remove(queue.Front())
		j = queue.Front().Value.(int)
		queue.Remove(queue.Front())

		addQueue(queue, board, i-1, j)
		addQueue(queue, board, i+1, j)
		addQueue(queue, board, i, j-1)
		addQueue(queue, board, i, j+1)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}
