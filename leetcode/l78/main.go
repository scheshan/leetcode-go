package l78

func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	res := make([][]int, 0)

	perm(nums, -1, &path, &res)
	return res
}

func perm(nums []int, ind int, path *[]int, res *[][]int) {
	r := make([]int, len(*path))
	copy(r, *path)
	*res = append(*res, r)

	for i := ind + 1; i < len(nums); i++ {
		*path = append(*path, nums[i])
		perm(nums, i, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
