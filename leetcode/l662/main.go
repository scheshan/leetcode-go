package l662

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type State struct {
	min int
	max int
}

func widthOfBinaryTree(root *TreeNode) int {
	m := make(map[int]*State)
	dfs(root, 1, 1, m)

	res := 0
	for _, state := range m {
		n := state.max - state.min + 1
		if n > res {
			res = n
		}
	}
	return res
}

func dfs(node *TreeNode, n int, level int, m map[int]*State) {
	if node == nil {
		return
	}

	state := m[level]
	if state == nil {
		state = &State{min: n, max: n}
		m[level] = state
	} else {
		if n < state.min {
			state.min = n
		}
		if n > state.max {
			state.max = n
		}
	}

	dfs(node.Left, n<<1, level+1, m)
	dfs(node.Right, (n<<1)+1, level+1, m)
}
