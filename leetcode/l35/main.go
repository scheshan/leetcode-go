package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)

	for l < r {
		mid := l + (r-l)>>1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if l >= 0 && l < len(nums) && nums[l] < target {
		return l + 1
	}

	return l
}

func main() {
	searchInsert([]int{1, 3, 5, 6}, 7)
}
