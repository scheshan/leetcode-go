package p087

import "fmt"

func restoreIpAddresses(s string) []string {
	path := make([]int, 0, 4)
	res := make([]string, 0)

	var backtrack func(int)
	backtrack = func(ind int) {
		if len(path) == 4 {
			if ind == len(s) {
				res = append(res, fmt.Sprintf("%v.%v.%v.%v", path[0], path[1], path[2], path[3]))
			}
			return
		}
		if ind >= len(s) {
			return
		}

		num := 0
		for i := ind; i < ind+3 && i < len(s); i++ {
			if i > ind && s[ind] == '0' {
				break
			}
			num = num*10 + int(s[i]-'0')
			if num > 255 {
				break
			}
			path = append(path, num)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
