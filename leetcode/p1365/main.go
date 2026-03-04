package p1365

func smallerNumbersThanCurrent(nums []int) []int {
	cnt := make([]int, 101)
	for _, num := range nums {
		cnt[num]++
	}

	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res[i] = 0
		} else {
			res[i] = cnt[nums[i]-1]
		}
	}

	return res
}
