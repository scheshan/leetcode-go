package j_05

import "strings"

func replaceSpace(s string) string {
	b := strings.Builder{}

	for _, ch := range s {
		if ch == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteRune(ch)
		}
	}

	return b.String()
}
