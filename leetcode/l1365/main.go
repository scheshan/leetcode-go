package l1365

func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)
	for _, n := range nums {
		count[n]++
	}

	pre := 0
	m := make(map[int]int, 0)
	for i, n := range count {
		if n > 0 {
			m[i] = pre
			pre += n
		}
	}

	res := make([]int, len(nums))
	for i, n := range nums {
		res[i] = m[n]
	}
	return res
}
