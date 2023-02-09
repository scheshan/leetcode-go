package p1258

import (
	"sort"
	"strings"
)

func generateSentences(synonyms [][]string, text string) []string {
	uf := make(map[string]string)
	var addWord = func(word string) {
		if _, ok := uf[word]; !ok {
			uf[word] = word
		}
	}
	var getParent = func(word string) string {
		if _, ok := uf[word]; !ok {
			return word
		}

		for uf[word] != word {
			uf[word] = uf[uf[word]]
			word = uf[word]
		}
		return uf[word]
	}
	var connect = func(s1 string, s2 string) {
		addWord(s1)
		addWord(s2)

		p1 := getParent(s1)
		p2 := getParent(s2)
		if p1 != p2 {
			uf[p1] = p2
		}
	}
	for _, arr := range synonyms {
		connect(arr[0], arr[1])
	}

	var buildWords = func() []string {
		res := make([]string, 0)
		pre := 0
		ind := 0
		for ind < len(text) {
			for ind < len(text) && text[ind] != ' ' {
				ind++
			}
			res = append(res, text[pre:ind])
			ind++
			pre = ind
		}
		return res
	}
	path := buildWords()
	res := make([]string, 0)
	var backtrack func(int)
	backtrack = func(ind int) {
		if ind == len(path) {
			res = append(res, strings.Join(path, " "))
			return
		}

		word := path[ind]
		backtrack(ind + 1)

		for next, _ := range uf {
			if word != next && getParent(word) == getParent(next) {
				path[ind] = next
				backtrack(ind + 1)
				path[ind] = word
			}
		}
	}
	backtrack(0)

	sort.Sort(sort.StringSlice(res))

	return res
}
