package l1315

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func sumEvenGrandparent(root *TreeNode) int {
	return travel(root, 1, 1)
}

func travel(node *TreeNode, grand int, parent int) int {
	if node == nil {
		return 0
	}

	n := travel(node.Left, parent, node.Val) + travel(node.Right, parent, node.Val)
	if grand&1 == 0 {
		n += node.Val
	}
	return n
}
