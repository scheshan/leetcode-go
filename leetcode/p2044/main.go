package p2044

func countMaxOrSubsets(nums []int) int {
	max := 0
	cnt := 0

	var backtrack func(int, int)
	backtrack = func(ind int, cur int) {
		if cur > max {
			max = cur
			cnt = 1
		} else if cur == max {
			cnt++
		}

		if ind == len(nums) {
			return
		}

		for i := ind; i < len(nums); i++ {
			backtrack(i+1, cur|nums[i])
		}
	}
	backtrack(0, 0)
	return cnt
}
