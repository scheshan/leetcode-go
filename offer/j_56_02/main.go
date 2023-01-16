package j_56_02

func singleNumber(nums []int) int {
	arr := make([]int, 26)

	for _, num := range nums {
		n := uint32(num)
		for i := 0; i < 32; i++ {
			if n&1 == 1 {
				arr[i]++
			}
			n >>= 1
		}
	}

	var res uint32
	for i := 31; i >= 0; i-- {
		res <<= 1
		if arr[i]%3 != 0 {
			res |= 1
		}
	}

	return int(res)
}
