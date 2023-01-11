package l1493

func longestSubarray(nums []int) int {
	res := 0
	flag := true

	l := 0
	r := 0

	for r < len(nums) {
		if nums[r] == 1 || flag {
			if nums[r] == 0 {
				flag = false
			}
			if r-l > res {
				res = r - l
			}
			r++
		} else {
			for !flag {
				if nums[l] == 0 {
					flag = true
				}
				l++
			}
		}
	}
	return res
}
