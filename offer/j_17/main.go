package j_17

func printNumbers(n int) []int {
	cur := 1
	next := cur * 10

	res := make([]int, 0)

	for i := 1; i <= n; i++ {
		for j := cur; j < next; j++ {
			res = append(res, j)
		}

		cur = next
		next = next * 10
	}

	return res
}
