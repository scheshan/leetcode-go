package l2500

import "sort"

type IntArray []int

func (t IntArray) Len() int {
	return len(t)
}

func (t IntArray) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t IntArray) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func deleteGreatestValue(grid [][]int) int {
	for _, row := range grid {
		ia := IntArray(row)
		sort.Sort(ia)
	}

	res := 0

	cols := len(grid[0])
	for col := 0; col < cols; col++ {
		max := 0
		for row := 0; row < len(grid); row++ {
			if grid[row][col] > max {
				max = grid[row][col]
			}
		}
		res += max
	}

	return res
}
