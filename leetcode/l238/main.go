package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		n := 1
		if i > 0 {
			n = res[i-1]
		}
		res[i] = nums[i] * n
	}

	cur := 1
	for i := len(nums) - 1; i >= 0; i-- {
		n := 1
		if i > 0 {
			n *= res[i-1]
		}
		res[i] = n * cur
		cur *= nums[i]
	}
	return res
}

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}
