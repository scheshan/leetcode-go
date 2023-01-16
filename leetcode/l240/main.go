package l240

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	n := 0

	for m >= 0 && n < len(matrix[m]) {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] < target {
			n++
		} else {
			m--
		}
	}

	return false
}
