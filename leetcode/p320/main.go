package p320

func generateAbbreviations(word string) []string {
	path := make([]byte, 0, len(word))
	res := make([]string, 0)

	var backtrack func(int, bool)
	backtrack = func(ind int, flag bool) {
		if ind == len(word) {
			res = append(res, string(path))
			return
		}

		if flag {
			for i := len(word) - 1; i >= ind; i-- {
				cnt := i - ind + 1
				if cnt < 10 {
					path = append(path, byte(cnt+'0'))
				} else {
					path = append(path, '1')
					path = append(path, byte(cnt-10+'0'))
				}
				backtrack(i+1, false)
				if cnt < 10 {
					path = path[:len(path)-1]
				} else {
					path = path[:len(path)-2]
				}
			}
		}

		path = append(path, word[ind])
		backtrack(ind+1, true)
		path = path[:len(path)-1]
	}
	backtrack(0, true)
	return res
}
