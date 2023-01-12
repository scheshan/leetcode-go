package main

import (
	"fmt"
	"sort"
)

type Arr []string

func (t Arr) Len() int {
	return len(t)
}

func (t Arr) Less(i, j int) bool {
	for k := 0; k < 5; k++ {
		if t[i][k] < t[j][k] {
			return true
		} else if t[i][k] > t[j][k] {
			return false
		}
	}
	return true
}

func (t Arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func findMinDifference(timePoints []string) int {
	arr := Arr(timePoints)
	sort.Sort(arr)

	res := 1440 - sum(timePoints[0], timePoints[len(timePoints)-1])
	for i := 1; i < len(timePoints); i++ {
		n := sum(timePoints[i-1], timePoints[i])
		if n < res {
			res = n
		}
	}
	return res
}

func sum(x, y string) int {
	hours := int(y[0])*10 + int(y[1]) - int(x[0])*10 - int(x[1])
	minutes := int(y[3])*10 + int(y[4]) - int(x[3])*10 - int(x[4])

	return hours*60 + minutes
}

func main() {
	difference := findMinDifference([]string{"00:00", "23:59"})
	fmt.Print(difference)
}
