package main

func singleNumbers(nums []int) []int {
	n := 0
	for _, num := range nums {
		n ^= num
	}

	x := 1
	for (n & 1) == 0 {
		x = x << 1
		n = n >> 1
	}

	a := 0
	b := 0
	for _, num := range nums {
		if num&x == x {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

func main() {
	singleNumbers([]int{4, 1, 4, 6})
}
