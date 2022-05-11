package l1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for j, num := range nums {
		delta := target - num

		if i, ok := m[delta]; ok {
			return []int{i, j}
		}

		m[num] = j
	}

	return nil
}
