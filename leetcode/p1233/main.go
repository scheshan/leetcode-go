package p1233

type Node struct {
	folder   string
	end      bool
	children map[string]*Node
}

func removeSubfolders(folder []string) []string {
	root := &Node{}
	for _, f := range folder {
		insert(root, f)
	}

	res := make([]string, 0)
	dfs(root, &res)
	return res
}

func insert(node *Node, folder string) {
	pre := 1
	ind := 1

	for ind < len(folder) {
		for ind < len(folder) && folder[ind] != '/' {
			ind++
		}

		seg := folder[pre:ind]
		if node.children == nil {
			node.children = make(map[string]*Node)
		}
		child := node.children[seg]
		if child == nil {
			child = &Node{}
			node.children[seg] = child
		}

		ind++
		pre = ind
		node = child
	}
	node.end = true
	node.folder = folder
}

func dfs(node *Node, res *[]string) {
	if node == nil {
		return
	}

	if node.end {
		*res = append(*res, node.folder)
		return
	}

	for _, child := range node.children {
		dfs(child, res)
	}
}
