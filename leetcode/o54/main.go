package o54

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type state struct {
	total int
}

func convertBST(root *TreeNode) *TreeNode {
	s := &state{}
	return travel(root, s)
}

func travel(root *TreeNode, state *state) *TreeNode {
	if root == nil {
		return nil
	}

	node := &TreeNode{}
	node.Right = travel(root.Right, state)
	node.Val = root.Val + state.total
	state.total += root.Val
	node.Left = travel(root.Left, state)
	return node
}
