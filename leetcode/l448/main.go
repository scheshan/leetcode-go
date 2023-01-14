package l448

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		if nums[num-1] > 0 {
			nums[num-1] *= -1
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
