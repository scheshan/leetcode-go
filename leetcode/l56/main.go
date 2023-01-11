package l56

import "sort"

type Arr [][]int

func (t Arr) Len() int {
	return len(t)
}

func (t Arr) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

func (t Arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func merge(intervals [][]int) [][]int {
	arr := Arr(intervals)
	sort.Sort(arr)

	res := make([][]int, 0)
	for _, pair := range arr {
		if len(res) == 0 {
			res = append(res, pair)
		} else {
			last := res[len(res)-1]
			if pair[0] <= last[1] {
				if pair[1] > last[1] {
					last[1] = pair[1]
				}
			} else {
				res = append(res, pair)
			}
		}
	}
	return res
}
