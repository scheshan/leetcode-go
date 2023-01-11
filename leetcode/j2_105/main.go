package j2_105

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				cnt := dfs(grid, i, j)
				if cnt > res {
					res = cnt
				}
			}
		}
	}

	return res
}

func dfs(grid [][]int, i int, j int) int {
	if grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = -1
	res := 1

	if i > 0 {
		res += dfs(grid, i-1, j)
	}
	if i < len(grid)-1 {
		res += dfs(grid, i+1, j)
	}
	if j > 0 {
		res += dfs(grid, i, j-1)
	}
	if j < len(grid[i])-1 {
		res += dfs(grid, i, j+1)
	}

	return res
}
