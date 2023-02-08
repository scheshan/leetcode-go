package main

func grayCode(n int) []int {
	total := 1 << n
	res := make([]int, 1, total)

	for i := 1; i <= n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]+(1<<(i-1)))
		}
	}
	return res
}
