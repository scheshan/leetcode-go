package l525

func findMaxLength(nums []int) int {
	m := make(map[int]int)

	n0 := 0
	res := 0
	for i, num := range nums {
		if num == 0 {
			n0++
		} else {
			n0--
		}

		if n0 == 0 && i+1 > res && (i+1)&1 == 0 {
			res = i + 1
		}

		if target, ok := m[n0]; ok && i-target > res {
			res = i - target
		}

		if _, ok := m[n0]; !ok {
			m[n0] = i
		}
	}

	return res
}
