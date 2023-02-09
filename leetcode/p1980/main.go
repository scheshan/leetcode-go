package p1980

func findDifferentBinaryString(nums []string) string {
	exist := make(map[string]bool)
	for _, num := range nums {
		exist[num] = true
	}

	path := make([]byte, len(nums))
	res := ""

	var backtrack func(ind int)
	backtrack = func(ind int) {
		if ind == len(path) {
			str := string(path)
			if !exist[str] {
				res = str
			}
			return
		}

		if res != "" {
			return
		}
		path[ind] = '0'
		backtrack(ind + 1)

		if res != "" {
			return
		}
		path[ind] = '1'
		backtrack(ind + 1)
	}
	backtrack(0)

	return res
}
