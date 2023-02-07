package p167

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		num := numbers[l] + numbers[r]
		if num == target {
			return []int{l + 1, r + 1}
		} else if num < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
