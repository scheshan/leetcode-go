package p169

func majorityElement(nums []int) int {
	res := 0
	vote := 0

	for i := 0; i < len(nums); i++ {
		if vote == 0 || nums[i] == res {
			res = nums[i]
			vote++
		} else {
			vote--
		}
	}

	return res
}
