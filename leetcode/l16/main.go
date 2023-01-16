package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	slice := sort.IntSlice(nums)
	sort.Sort(slice)

	res := math.MaxInt
	for ind, num := range nums {
		if ind > 0 && nums[ind] == nums[ind-1] {
			continue
		}

		l := ind + 1
		r := len(nums) - 1

		for l < r {
			n := nums[l] + num + nums[r]
			if n == target {
				return target
			}
			if abs(target-n) < abs(target-res) {
				res = n
			}

			if n < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	threeSumClosest([]int{-100, -98, -2, -1}, -101)
}
