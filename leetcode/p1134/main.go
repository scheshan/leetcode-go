package p1134

import (
	"math"
	"strconv"
)

func isArmstrong(n int) bool {
	str := strconv.Itoa(n)

	s := 0
	for _, ch := range str {
		pow := math.Pow(float64(ch-'0'), float64(len(str)))
		s += int(pow)
	}

	return s == n
}
