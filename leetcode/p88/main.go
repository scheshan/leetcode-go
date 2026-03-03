package p88

func merge(nums1 []int, m int, nums2 []int, n int) {
	ind := m + n - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[ind] = nums1[m-1]
			m--
		} else {
			nums1[ind] = nums2[n-1]
			n--
		}
		ind--
	}

	for n > 0 {
		nums1[ind] = nums2[n-1]
		n--
		ind--
	}
}
