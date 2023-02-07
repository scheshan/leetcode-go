package p38

import "sort"

type ByteSlice []byte

func (t ByteSlice) Len() int {
	return len(t)
}

func (t ByteSlice) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t ByteSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func permutation(s string) []string {
	arr := []byte(s)
	sort.Sort(ByteSlice(arr))
	res := make([]string, 0)
	path := make([]byte, 0, len(arr))
	visit := make([]bool, len(arr))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(s) {
			res = append(res, string(path))
			return
		}

		for i := 0; i < len(arr); i++ {
			if visit[i] {
				continue
			}
			if i > 0 && arr[i] == arr[i-1] && !visit[i-1] {
				continue
			}
			path = append(path, arr[i])
			visit[i] = true
			backtrack()
			visit[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return res
}
