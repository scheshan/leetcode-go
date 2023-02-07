package p08_07

func permutation(S string) []string {
	arr := []byte(S)
	res := make([]string, 0)

	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == len(arr) {
			res = append(res, string(arr))
			return
		}

		for i := ind; i < len(arr); i++ {
			arr[i], arr[ind] = arr[ind], arr[i]
			backtrack(ind + 1)
			arr[i], arr[ind] = arr[ind], arr[i]
		}
	}
	backtrack(0)
	return res
}
