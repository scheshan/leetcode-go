package main

import "fmt"

type Trie struct {
	root *trieNode
}

type trieNode struct {
	end      bool
	children []*trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{
		root: &trieNode{},
	}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for i, ch := range word {
		if node.children == nil {
			node.children = make([]*trieNode, 26)
		}

		child := node.children[ch-'a']
		if child == nil {
			child = &trieNode{}
			node.children[ch-'a'] = child
		}

		node = child
		if i == len(word)-1 {
			child.end = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for i, ch := range word {
		if node.children == nil {
			return false
		}

		child := node.children[ch-'a']
		if child == nil {
			return false
		}

		if i == len(word)-1 {
			return child.end
		}
		node = child
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		if node.children == nil {
			return false
		}

		child := node.children[ch-'a']
		if child == nil {
			return false
		}
		node = child
	}
	return true
}

func main() {
	t := Constructor()
	trie := &t

	trie.Insert("apple")
	fmt.Print(trie.Search("apple"))
}
