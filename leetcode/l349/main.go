package l349

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)

	for _, num := range nums1 {
		m[num] = true
	}

	res := make([]int, 0, len(nums2))

	for _, num := range nums2 {
		if m[num] {
			res = append(res, num)
			m[num] = false
		}
	}

	return res
}
