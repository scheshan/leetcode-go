package l1748

func sumOfUnique(nums []int) int {
	m := make(map[int]int)
	res := 0

	for _, num := range nums {
		cnt := m[num]
		if cnt == 0 {
			res += num
		} else if cnt == 1 {
			res -= num
		}
		m[num] = cnt + 1
	}
	return res
}
