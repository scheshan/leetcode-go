package p283

func moveZeroes(nums []int) {
	slow := 0
	fast := 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}