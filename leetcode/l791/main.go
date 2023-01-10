package l791

import "strings"

func customSortString(order string, s string) string {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a'] = cnt[ch-'a'] + 1
	}

	builder := strings.Builder{}
	for _, ch := range order {
		if cnt[ch-'a'] > 0 {
			for i := 0; i < cnt[ch-'a']; i++ {
				builder.WriteRune(ch)
			}
			cnt[ch-'a'] = 0
		}
	}
	for i, n := range cnt {
		for j := 0; j < n; j++ {
			builder.WriteRune('a' + rune(i))
		}
	}

	return builder.String()
}
