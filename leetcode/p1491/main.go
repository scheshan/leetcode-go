package p1491

import "math"

func average(salary []int) float64 {
	min := math.MaxInt
	max := math.MinInt
	total := 0

	for _, num := range salary {
		total += num
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return float64(total-min-max) / float64(len(salary)-2)
}
