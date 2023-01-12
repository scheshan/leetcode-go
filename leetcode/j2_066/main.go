package main

type MapSum struct {
	root *Node
}

type Node struct {
	count    int
	children []*Node
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	m := MapSum{
		root: &Node{},
	}
	return m
}

func (this *MapSum) Insert(key string, val int) {
	node := this.root
	for ind, ch := range key {
		if node.children == nil {
			node.children = make([]*Node, 26)
		}
		child := node.children[ch-'a']
		if child == nil {
			child = &Node{}
			node.children[ch-'a'] = child
		}

		if ind == len(key)-1 {
			child.count = val
		}
		node = child
	}
}

func (this *MapSum) Sum(prefix string) int {
	node := this.root
	for _, ch := range prefix {
		if node.children == nil || node.children[ch-'a'] == nil {
			return 0
		}
		node = node.children[ch-'a']
	}
	return this.dfs(node)
}

func (t *MapSum) dfs(node *Node) int {
	if node == nil {
		return 0
	}

	res := node.count
	for _, child := range node.children {
		if child != nil {
			res += t.dfs(child)
		}
	}
	return res
}

func main() {
	ms := Constructor()
	ms.Insert("apple", 3)
	println(ms.Sum("ap"))
	ms.Insert("app", 2)
	println(ms.Sum("ap"))
}
