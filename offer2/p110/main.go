package p110

func allPathsSourceTarget(graph [][]int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(ind int)
	dfs = func(ind int) {
		path = append(path, ind)
		if ind == len(graph)-1 {
			arr := make([]int, len(path))
			copy(arr, path)
			res = append(res, arr)
		} else {
			edge := graph[ind]
			for _, next := range edge {
				dfs(next)
			}
		}
		path = path[:len(path)-1]
	}
	dfs(0)
	return res
}
