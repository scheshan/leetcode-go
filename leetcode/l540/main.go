package l540

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}

	l := 0
	r := len(nums)

	res := 0
	for l < r {
		mid := l + (r-l)>>1

		if mid&1 == 0 {
			if nums[mid] != nums[mid-1] {
				l = mid + 1
			} else {
				r = mid
				res = nums[mid-1]
			}
		} else {
			if nums[mid] == nums[mid-1] {
				l = mid + 1
			} else {
				r = mid
				res = nums[mid-1]
			}
		}
	}
	return res
}
