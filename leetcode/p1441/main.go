package p1441

func buildArray(target []int, n int) []string {
	ind := 0
	cur := 0
	res := make([]string, 0)

	for ind < len(target) {
		cur++
		res = append(res, "Push")
		if cur < target[ind] {
			res = append(res, "Pop")
		} else {
			ind++
		}
	}

	return res
}
