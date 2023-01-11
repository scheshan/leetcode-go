package l797

func allPathsSourceTarget(graph [][]int) [][]int {
	path := make([]int, 0, len(graph))
	path = append(path, 0)
	visit := make([]bool, len(graph))
	visit[0] = true
	res := make([][]int, 0)

	dfs(graph, 0, visit, &path, &res)

	return res
}

func dfs(graph [][]int, n int, visit []bool, path *[]int, res *[][]int) {
	if n == len(graph)-1 {
		r := make([]int, len(*path))
		copy(r, *path)
		*res = append(*res, r)
		return
	}

	for _, next := range graph[n] {
		if visit[next] {
			continue
		}

		visit[next] = true
		*path = append(*path, next)

		dfs(graph, next, visit, path, res)

		visit[next] = false
		*path = (*path)[:len(*path)-1]
	}
}
