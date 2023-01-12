package l648

import "strings"

type Node struct {
	end      bool
	children map[rune]*Node
}

func replaceWords(dictionary []string, sentence string) string {
	root := &Node{}
	for _, word := range dictionary {
		insert(root, word)
	}

	b := strings.Builder{}

	l := 0
	r := 0
	for r < len(sentence) {
		for r < len(sentence) && sentence[r] != ' ' {
			r++
		}

		n := find(root, sentence, l, r)
		b.Write([]byte(sentence[l : l+n]))
		if r < len(sentence) {
			r++
			b.WriteRune(' ')
		}
		l = r
	}
	return b.String()
}

func insert(root *Node, word string) {
	node := root

	for i, ch := range word {
		if node.children == nil {
			node.children = make(map[rune]*Node)
		}

		child := node.children[ch-'a']
		if child == nil {
			child = &Node{}
			node.children[ch-'a'] = child
		}
		if i == len(word)-1 {
			child.end = true
		}

		node = child
	}
}

func find(root *Node, sentence string, begin int, end int) int {
	node := root

	ind := begin
	for ind < end {
		if node.children == nil {
			return end - begin
		}
		node = node.children[rune(sentence[ind]-'a')]
		if node == nil {
			return end - begin
		}
		if node.end {
			return ind - begin + 1
		}
		ind++
	}

	return end - begin
}
