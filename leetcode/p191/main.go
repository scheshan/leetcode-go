package p191

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num = num & (num - 1)
	}
	return res
}
