package j_53_01

func search(nums []int, target int) int {
	return findLarger(nums, target) - findLess(nums, target) - 1
}

func findLess(nums []int, target int) int {
	l := 0
	r := len(nums)
	res := -1

	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			res = mid
			l = mid + 1
		} else {
			r = mid
		}
	}

	return res
}

func findLarger(nums []int, target int) int {
	l := 0
	r := len(nums)
	res := len(nums)

	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			res = mid
			r = mid
		} else {
			l = mid + 1
		}
	}

	return res
}
