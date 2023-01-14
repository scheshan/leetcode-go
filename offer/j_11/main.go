package main

func minArray(numbers []int) int {
	l := 0
	r := len(numbers) - 1

	for l <= r {
		mid := l + (r-l)>>1
		n := numbers[mid]

		if n > numbers[r] {
			l = mid + 1
		} else if n < numbers[r] {
			r = mid
		} else {
			r--
		}
	}

	return numbers[l]
}
