package l46

func permute(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)

	perm(nums, make(map[int]bool), &path, &res)
	return res
}

func perm(nums []int, visit map[int]bool, path *[]int, res *[][]int) {
	if len(*path) == len(nums) {
		r := make([]int, len(nums), len(nums))
		copy(r, *path)
		*res = append(*res, r)
		return
	}

	for _, num := range nums {
		if !visit[num] {
			visit[num] = true
			*path = append(*path, num)
			perm(nums, visit, path, res)
			*path = (*path)[:len(*path)-1]
			visit[num] = false
		}
	}
}
