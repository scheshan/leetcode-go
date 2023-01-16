package l154

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)>>1
		n := nums[mid]

		if n > nums[r] {
			l = mid + 1
		} else if n < nums[r] {
			r = mid
		} else {
			r--
		}
	}

	return nums[l]
}
