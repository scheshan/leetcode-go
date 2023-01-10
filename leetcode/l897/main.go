package l897

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type state struct {
	head *TreeNode
	tail *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	s := &state{
		head: &TreeNode{},
	}
	s.tail = s.head
	travel(root, s)
	return s.head.Right
}

func travel(root *TreeNode, state *state) {
	if root == nil {
		return
	}

	travel(root.Left, state)

	node := &TreeNode{Val: root.Val}
	state.tail.Right = node
	state.tail = node

	travel(root.Right, state)
}
