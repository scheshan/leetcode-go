package p1238

func circularPermutation(n int, start int) []int {
	arr := make([]int, 1, 1<<n)

	ind := 0
	for i := 1; i <= n; i++ {
		for j := len(arr) - 1; j >= 0; j-- {
			arr = append(arr, arr[j]+(1<<(i-1)))
			if arr[len(arr)-1] == start {
				ind = len(arr) - 1
			}
		}
	}

	res := make([]int, 0, len(arr))
	for i := ind; i < len(arr); i++ {
		res = append(res, arr[i])
	}
	for i := 0; i < ind; i++ {
		res = append(res, arr[i])
	}
	return res
}
