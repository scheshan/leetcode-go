package l230

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	s := &state{count: k}
	dfs(root, s)
	return s.num
}

type state struct {
	num   int
	count int
}

func dfs(root *TreeNode, state *state) {
	if root == nil || state.count == 0 {
		return
	}

	dfs(root.Left, state)
	if state.count > 0 {
		state.count--

		if state.count == 0 {
			state.num = root.Val
		}
	}
	dfs(root.Right, state)
}
