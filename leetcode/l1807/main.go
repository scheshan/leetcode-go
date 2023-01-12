package l1807

import "strings"

func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	for _, k := range knowledge {
		m[k[0]] = k[1]
	}

	b := strings.Builder{}
	ind := 0
	for ind < len(s) {
		if s[ind] == '(' {
			r := ind + 1
			for s[r] != ')' {
				r++
			}

			str := m[s[ind+1:r]]
			if str == "" {
				str = "?"
			}
			b.WriteString(str)
			ind = r + 1
		} else {
			b.WriteByte(s[ind])
			ind++
		}
	}
	return b.String()
}
