package p77

func combine(n int, k int) [][]int {
	path := make([]int, 0, k)
	res := make([][]int, 0)
	backtrack(n, 1, k, &path, &res)
	return res
}

func backtrack(n int, ind int, k int, path *[]int, res *[][]int) {
	if len(*path) == k {
		arr := make([]int, k)
		copy(arr, *path)
		*res = append(*res, arr)
		return
	}
	if ind > n {
		return
	}

	for i := ind; i <= n; i++ {
		*path = append(*path, i)
		backtrack(n, i+1, k, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
