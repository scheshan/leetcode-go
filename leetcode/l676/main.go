package l676

type MagicDictionary struct {
	root *Node
}

type Node struct {
	end      bool
	children []*Node
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	dict := MagicDictionary{
		root: &Node{},
	}
	return dict
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.insert(word)
	}
}

func (t *MagicDictionary) insert(word string) {
	node := t.root

	for ind, ch := range word {
		if node.children == nil {
			node.children = make([]*Node, 26)
		}

		child := node.children[ch-'a']
		if child == nil {
			child = &Node{}
			node.children[ch-'a'] = child
		}
		if ind == len(word)-1 {
			child.end = true
		}
		node = child
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.find(this.root, searchWord, 0, false)
}

func (t *MagicDictionary) find(node *Node, word string, ind int, used bool) bool {
	if ind == len(word) {
		return node.end && used
	}
	if node.children == nil {
		return false
	}

	ch := word[ind]
	child := node.children[ch-'a']

	if child != nil && t.find(child, word, ind+1, used) {
		return true
	}

	if !used {
		for i := 0; i < len(node.children); i++ {
			if node.children[i] != nil && i != int(ch-'a') && t.find(node.children[i], word, ind+1, true) {
				return true
			}
		}
	}

	return false
}
