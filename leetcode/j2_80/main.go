package j2_80

func combine(n int, k int) [][]int {
	path := make([]int, 0, k)
	res := make([][]int, 0)
	perm(n, k, 0, &path, &res)

	return res
}

func perm(n int, k int, cur int, path *[]int, res *[][]int) {
	if len(*path) == k {
		r := make([]int, k)
		copy(r, *path)
		*res = append(*res, r)
	} else {
		for i := cur + 1; i <= n; i++ {
			*path = append(*path, i)
			perm(n, k, i, path, res)
			*path = (*path)[:len(*path)-1]
		}
	}
}
