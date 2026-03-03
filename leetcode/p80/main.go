package p80

func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0

	for fast < len(nums) {
		num := nums[fast]
		nums[slow] = num
		fast++

		for slow > 0 && fast < len(nums) && nums[fast] == nums[slow-1] {
			fast++
		}

		slow++
	}

	return slow
}
