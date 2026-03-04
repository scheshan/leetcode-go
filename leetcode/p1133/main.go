package p1133

func largestUniqueNumber(nums []int) int {
	cnt := make(map[int]int)

	for _, n := range nums {
		cnt[n]++
	}

	res := -1
	for k, v := range cnt {
		if v == 1 {
			res = max(res, k)
		}
	}

	return res
}
