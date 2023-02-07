package p17

func printNumbers(n int) []int {
	size := 1
	for i := 0; i < n; i++ {
		size *= 10
	}
	size--

	res := make([]int, size)
	for i := 1; i <= size; i++ {
		res[i-1] = i
	}
	return res
}
