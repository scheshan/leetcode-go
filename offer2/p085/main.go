package p085

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0, n<<1)
	backtrack(n, 0, 0, &path, &res)
	return res
}

func backtrack(n int, left int, right int, path *[]byte, res *[]string) {
	if len(*path) == n<<1 {
		*res = append(*res, string(*path))
		return
	}

	if left < n {
		*path = append(*path, '(')
		backtrack(n, left+1, right, path, res)
		*path = (*path)[:len(*path)-1]
	}
	if right < left {
		*path = append(*path, ')')
		backtrack(n, left, right+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
