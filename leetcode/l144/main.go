package l144

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type Stack struct {
	arr []*TreeNode
}

func (t *Stack) Push(node *TreeNode) {
	t.arr = append(t.arr, node)
}

func (t *Stack) Pop() *TreeNode {
	node := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return node
}

func (t *Stack) IsEmpty() bool {
	return len(t.arr) == 0
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	stack := &Stack{arr: make([]*TreeNode, 0)}
	stack.Push(root)

	for !stack.IsEmpty() {
		node := stack.Pop()
		if node != nil {
			if node.Right != nil {
				stack.Push(node.Right)
			}
			if node.Left != nil {
				stack.Push(node.Left)
			}
			stack.Push(node)
			stack.Push(nil)
		} else {
			res = append(res, stack.Pop().Val)
		}
	}
	return res
}
