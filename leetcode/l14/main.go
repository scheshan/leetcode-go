package l14

import "strings"

func longestCommonPrefix(strs []string) string {
	b := strings.Builder{}

	ind := 0
	for {
		var ch byte = ' '
		for i := 0; i < len(strs); i++ {
			if ind >= len(strs[i]) {
				return b.String()
			}
			if ch == ' ' {
				ch = strs[i][ind]
			} else if ch != strs[i][ind] {
				return b.String()
			}
		}
		b.WriteByte(ch)
		ind++
	}
	return b.String()
}
