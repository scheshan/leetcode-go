package l704

func search(nums []int, target int) int {
	l := 0
	r := len(nums)

	for l < r {
		mid := (l + r) >> 1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return -1
}
