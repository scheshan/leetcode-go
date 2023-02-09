package p816

import "fmt"

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	path := make([]string, 0, len(s))
	res := make([]string, 0)

	var backtrack func(ind int)
	backtrack = func(ind int) {
		if len(path) == 2 {
			if ind == len(s) {
				res = append(res, fmt.Sprintf("(%s, %s)", path[0], path[1]))
			}
			return
		}

		for i := ind; i < len(s); i++ {
			if i > ind {
				for j := ind; j < i; j++ {
					if s[i] == '0' {
						break
					}

					left := s[ind : j+1]
					if left[0] == '0' && len(left) > 1 {
						break
					}
					right := s[j+1 : i+1]
					path = append(path, left+"."+right)
					backtrack(i + 1)
					path = path[:len(path)-1]
				}
			}

			if i == ind || s[ind] != '0' {
				path = append(path, s[ind:i+1])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return res
}
