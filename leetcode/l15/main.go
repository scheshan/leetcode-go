package main

import "sort"

func threeSum(nums []int) [][]int {
	slice := sort.IntSlice(nums)
	sort.Sort(slice)

	res := make([][]int, 0)

	for ind, num := range nums {
		if num > 0 {
			return res
		}
		if ind > 0 && nums[ind] == nums[ind-1] {
			continue
		}

		l := ind + 1
		r := len(nums) - 1
		for l < r {
			n := nums[l] + num + nums[r]
			if n == 0 {
				res = append(res, []int{num, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else if n < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
