package l2089

func targetIndices(nums []int, target int) []int {
	n1 := 0
	n2 := 0
	for _, num := range nums {
		if num < target {
			n1++
		} else if num == target {
			n2++
		}
	}

	res := make([]int, n2)
	for i := 0; i < n2; i++ {
		res[i] = n1 + i
	}
	return res
}
