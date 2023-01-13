package j_53_02

func missingNumber(nums []int) int {
	l := 0
	r := len(nums)

	res := len(nums)
	for l < r {
		mid := l + (r-l)>>1

		if nums[mid] > mid {
			r = mid
			res = mid
		} else {
			l = mid + 1
		}
	}

	return res
}
