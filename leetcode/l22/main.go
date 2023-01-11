package l22

func generateParenthesis(n int) []string {
	cache := make([]rune, n<<1)
	res := make([]string, 0)

	perm(n, 0, 0, &cache, &res)

	return res
}

func perm(n int, left int, right int, cache *[]rune, res *[]string) {
	if left > n || right > n || right > left {
		return
	}
	if left == n && right == n {
		str := string(*cache)
		*res = append(*res, str)
		return
	}

	ind := left + right
	(*cache)[ind] = '('
	perm(n, left+1, right, cache, res)

	(*cache)[ind] = ')'
	perm(n, left, right+1, cache, res)
}
