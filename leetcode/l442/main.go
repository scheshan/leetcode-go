package l442

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for _, n := range nums {
		n = abs(n)
		if nums[n-1] < 0 {
			res = append(res, n)
		}
		nums[n-1] = -1 * nums[n-1]
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
