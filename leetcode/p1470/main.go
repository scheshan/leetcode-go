package p1470

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	ind := 0
	for i := 0; i < n; i++ {
		res[ind] = nums[i]
		ind++
		res[ind] = nums[i+n]
		ind++
	}
	return res
}
