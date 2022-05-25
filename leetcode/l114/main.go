package l114

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root)

	cur := new(TreeNode)
	var node *TreeNode

	for stack.Len() > 0 {
		ele := stack.Back()
		stack.Remove(ele)

		node = ele.Value.(*TreeNode)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		node.Left = nil
		node.Right = nil
		cur.Right = node
		cur = node
	}
}
