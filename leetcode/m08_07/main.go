package m08_07

func permutation(S string) []string {
	res := make([]string, 0)
	perm(&res, []byte(S), 0)

	return res
}

func perm(res *[]string, arr []byte, ind int) {
	if ind == len(arr) {
		*res = append(*res, string(arr))
		return
	}

	for i := ind; i < len(arr); i++ {
		arr[ind], arr[i] = arr[i], arr[ind]

		perm(res, arr, ind+1)

		arr[ind], arr[i] = arr[i], arr[ind]
	}
}
