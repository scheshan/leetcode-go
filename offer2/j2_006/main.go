package j2_006

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		n := numbers[l] + numbers[r]
		if n == target {
			return []int{l, r}
		} else if n < target {
			l++
		} else {
			r--
		}
	}

	return nil
}
