package p645

func findErrorNums(nums []int) []int {
	dup := 0

	for i := 0; i < len(nums); i++ {
		ind := nums[i]
		if ind < 0 {
			ind = -ind
		}

		if nums[ind-1] < 0 {
			dup = ind
		} else {
			nums[ind-1] *= -1
		}
	}

	res := make([]int, 0, 2)
	res = append(res, dup)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
