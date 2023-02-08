package p1079

import (
	"sort"
)

func numTilePossibilities(tiles string) int {
	arr := make([]int, len(tiles))
	for i, ch := range tiles {
		arr[i] = int(ch - 'A')
	}
	sort.Sort(sort.IntSlice(arr))
	visit := make([]bool, len(arr))
	path := make([]byte, 0, len(arr))

	res := 0
	var backtrack func()
	backtrack = func() {
		if len(path) > 0 {
			res++
		}
		if len(path) == len(arr) {
			return
		}

		for i := 0; i < len(arr); i++ {
			if visit[i] {
				continue
			}
			if i > 0 && arr[i] == arr[i-1] && !visit[i-1] {
				continue
			}

			visit[i] = true
			path = append(path, byte(arr[i]+'A'))
			backtrack()
			visit[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return res
}
