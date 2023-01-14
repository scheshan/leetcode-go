package main

func singleNumber(nums []int) int {
	arr := make([]int, 32)

	for _, num := range nums {
		n := int32(num)
		for i := 0; i < 32; i++ {
			if n&1 != 0 {
				arr[i]++
			}
			n >>= 1
		}
	}

	res := int32(0)
	for i := 0; i < len(arr); i++ {
		if arr[i]%3 > 0 {
			res = res | (1 << i)
		}
	}

	return int(res)
}

func main() {
	singleNumber([]int{-2, -2, -2, -4})
}
