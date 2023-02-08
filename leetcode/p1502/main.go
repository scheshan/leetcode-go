package p1502

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Sort(sort.IntSlice(arr))

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != arr[1]-arr[0] {
			return false
		}
	}
	return true
}
