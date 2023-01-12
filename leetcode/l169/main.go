package l169

func majorityElement(nums []int) int {
	vote := 0
	res := 0

	for _, num := range nums {
		if vote == 0 {
			vote++
			res = num
		} else if num != res {
			vote--
		} else {
			vote++
		}
	}

	return res
}
