package j2_110

func allPathsSourceTarget(graph [][]int) [][]int {
	path := make([]int, 0)
	path = append(path, 0)
	res := make([][]int, 0)

	dfs(graph, 0, make(map[int]bool), &path, &res)

	return res
}

func dfs(graph [][]int, n int, visit map[int]bool, path *[]int, res *[][]int) {
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

		delete(visit, next)
		*path = (*path)[:len(*path)-1]
	}
}
