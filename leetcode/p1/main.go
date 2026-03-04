package p1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if pre, ok := m[target-n]; ok {
			return []int{pre, i}
		}
		m[n] = i
	}

	return nil
}
