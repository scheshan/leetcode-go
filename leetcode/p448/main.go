package p448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		ind := nums[i]
		if ind < 0 {
			ind *= -1
		}
		if nums[ind-1] > 0 {
			nums[ind-1] *= -1
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}
