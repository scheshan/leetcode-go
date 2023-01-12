package j_54

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type state struct {
	num   int
	count int
}

func kthLargest(root *TreeNode, k int) int {
	s := &state{count: k}
	dfs(root, s)
	return s.num
}

func dfs(root *TreeNode, state *state) {
	if root == nil || state.count == 0 {
		return
	}

	dfs(root.Right, state)
	if state.count > 0 {
		state.count--

		if state.count == 0 {
			state.num = root.Val
		}
	}
	dfs(root.Left, state)
}
