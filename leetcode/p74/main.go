package p74

func searchMatrix(matrix [][]int, target int) bool {
	row := -1
	top := 0
	down := len(matrix)

	for top < down {
		mid := top + ((down - top) >> 1)
		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target {
			row = mid
			break
		} else if matrix[mid][0] > target {
			down = mid
		} else {
			top = mid + 1
		}
	}

	if row == -1 {
		return false
	}

	l := 0
	r := len(matrix[row])

	for l < r {
		mid := l + ((r - l) >> 1)
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}
