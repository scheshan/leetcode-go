package l1004

func longestOnes(nums []int, k int) int {
	l := 0
	r := 0

	res := 0
	for r < len(nums) {
		if nums[r] == 1 || k > 0 {
			if nums[r] == 0 {
				k--
			}

			r++
			if r-l > res {
				res = r - l
			}
		} else {
			for k == 0 {
				if nums[l] == 0 {
					k++
				}
				l++
			}
		}
	}
	return res
}
