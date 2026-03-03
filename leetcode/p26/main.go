package p26

func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0

	for fast < len(nums) {
		nums[slow] = nums[fast]
		slow++
		fast++
		for fast < len(nums) && nums[fast] == nums[fast-1] {
			fast++
		}
	}

	return slow
}
