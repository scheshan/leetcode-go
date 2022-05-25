package l1295

import "math"

func findNumbers(nums []int) int {
	res := 0

	for _, n := range nums {
		if int(math.Log10(float64(n)))&1 == 0 {
			res++
		}
	}

	return res
}
