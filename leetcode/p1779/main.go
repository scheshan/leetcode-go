package p1779

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	min := math.MaxInt
	res := -1

	for i, p := range points {
		if p[0] == x || p[1] == y {
			dis := abs(p[0]-x) + abs(p[1]-y)
			if dis < min {
				min = dis
				res = i
			}
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
