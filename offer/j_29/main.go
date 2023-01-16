package main

type Point struct {
	x int
	y int
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0, len(matrix)*len(matrix[0]))

	tLeft := Point{x: 0, y: 0}
	tRight := Point{x: 0, y: len(matrix[0]) - 1}
	bLeft := Point{x: len(matrix) - 1, y: 0}
	bRight := Point{x: len(matrix) - 1, y: len(matrix[0]) - 1}

	for tLeft.x <= bRight.x && tLeft.y <= bRight.y {
		for y := tLeft.y; y <= tRight.y; y++ {
			res = append(res, matrix[tLeft.x][y])
		}
		for x := tRight.x + 1; x <= bRight.x; x++ {
			res = append(res, matrix[x][bRight.y])
		}
		if tLeft.x < bRight.x && tLeft.y < bRight.y {
			for y := bRight.y - 1; y >= bLeft.y; y-- {
				res = append(res, matrix[bRight.x][y])
			}
			for x := bLeft.x - 1; x >= tLeft.x+1; x-- {
				res = append(res, matrix[x][tLeft.y])
			}
		}

		tLeft.x++
		tLeft.y++
		tRight.x++
		tRight.y--
		bLeft.x--
		bLeft.y++
		bRight.x--
		bRight.y--
	}

	return res
}

func main() {
	spiralOrder([][]int{{1, 2, 3}})
}
