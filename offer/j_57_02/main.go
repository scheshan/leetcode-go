package main

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)

	l := 1
	r := 1
	cur := 0
	for r <= (target>>1)+1 {
		cur += r

		for cur > target {
			cur -= l
			l++
		}

		if cur == target {
			arr := make([]int, 0, r-l+1)
			for i := l; i <= r; i++ {
				arr = append(arr, i)
			}
			res = append(res, arr)
		}

		r++
	}

	return res
}

func main() {
	findContinuousSequence(6)
}
