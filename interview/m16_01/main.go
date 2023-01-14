package m16_01

func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] + numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] = numbers[0] - numbers[1]

	return numbers
}
