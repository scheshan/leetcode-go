package p1929

func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)<<1)
	for i := 0; i < len(nums); i++ {
		res[i] = nums[i]
		res[i+len(nums)] = nums[i]
	}

	return res
}
