package l343

func integerBreak(n int) int {
	res := 0
	m := make(map[int]int)

	for i := 1; i < n; i++ {
		res = max(res, dfs(i, m)*dfs(n-i, m))
	}

	return res
}

func dfs(n int, m map[int]int) int {
	if n <= 2 {
		return n
	}

	if res, ok := m[n]; ok {
		return res
	}

	res := n
	for i := 1; i < n; i++ {
		res = max(dfs(i, m)*dfs(n-i, m), res)
	}

	m[n] = res
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
