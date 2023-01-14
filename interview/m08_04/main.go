package main

func subsets(nums []int) [][]int {
	res := dfs(nums, 0)
	res = append(res, []int{})
	return res
}

func dfs(nums []int, ind int) [][]int {
	if ind >= len(nums) {
		return nil
	}

	res := dfs(nums, ind+1)
	size := len(res)
	for i := 0; i < size; i++ {
		items := make([]int, len(res[i]), len(res[i])+1)
		copy(items, res[i])
		items = append(items, nums[ind])

		res = append(res, items)
	}
	res = append(res, []int{nums[ind]})
	return res
}

func main() {
	subsets([]int{9, 0, 3, 5, 7})
}
