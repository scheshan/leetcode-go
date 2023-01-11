package main

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	dfs(candidates, 0, target, &path, &res)
	return res
}

func dfs(candidates []int, ind int, target int, path *[]int, res *[][]int) {
	if target == 0 {
		r := make([]int, len(*path))
		copy(r, *path)
		*res = append(*res, r)
		return
	}
	if target < 0 || ind >= len(candidates) {
		return
	}

	*path = append(*path, candidates[ind])
	dfs(candidates, ind, target-candidates[ind], path, res)
	*path = (*path)[:len(*path)-1]

	for i := ind + 1; i < len(candidates); i++ {
		num := candidates[i]
		*path = append(*path, num)
		dfs(candidates, i, target-num, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	combinationSum([]int{2, 3, 5}, 8)
}
