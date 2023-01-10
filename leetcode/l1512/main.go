package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num] = m[num] + 1
	}

	res := 0
	for _, n := range m {
		res += count(n)
	}
	return res
}

func count(n int) int {
	if n <= 1 {
		return 0
	}

	return n * (n - 1) / 2
}

func main() {
	res := numIdenticalPairs([]int{1, 1, 1, 1})
	fmt.Print(res)
}
