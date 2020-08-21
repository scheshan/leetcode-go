package l1313

func decompressRLElist(nums []int) []int {
	var res []int

	i := 0
	for i < len(nums) {
		for nums[i] > 0 {
			res = append(res, nums[i+1])
			nums[i] -= 1
		}
		i += 2
	}

	return res
}
