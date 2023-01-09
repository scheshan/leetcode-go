package l538

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type state struct {
	total int
}

func convertBST(root *TreeNode) *TreeNode {
	s := &state{}
	travel(root, s)

	return root
}

func travel(root *TreeNode, state *state) {
	if root == nil {
		return
	}

	travel(root.Right, state)

	state.total += root.Val
	root.Val = state.total

	travel(root.Left, state)
}
