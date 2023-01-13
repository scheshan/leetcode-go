package l34

func searchRange(nums []int, target int) []int {
	less := findLess(nums, target)
	larger := findLarger(nums, target)

	if larger-less-1 == 0 {
		return []int{-1, -1}
	}
	return []int{less + 1, larger - 1}
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
