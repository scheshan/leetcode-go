package l977

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	ind := 0
	for ind < len(nums) && nums[ind] < 0 {
		ind++
	}

	left := ind - 1
	right := ind

	ind = 0

	for left >= 0 && right < len(nums) {
		ln := abs(nums[left])
		rn := abs(nums[right])

		if ln < rn {
			res[ind] = ln * ln
			left--
		} else {
			res[ind] = rn * rn
			right++
		}
		ind++
	}

	for left >= 0 {
		res[ind] = nums[left] * nums[left]
		left--
		ind++
	}
	for right < len(nums) {
		res[ind] = nums[right] * nums[right]
		right++
		ind++
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
