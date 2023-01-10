package l804

import "strings"

func uniqueMorseRepresentations(words []string) int {
	keyMap := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	m := make(map[string]bool)

	for _, word := range words {
		b := strings.Builder{}

		for _, ch := range word {
			b.WriteString(keyMap[ch-'a'])
		}

		m[b.String()] = true
	}

	return len(m)
}
