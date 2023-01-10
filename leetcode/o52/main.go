package o52

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

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	dummy := &TreeNode{}
	tail := dummy

	stack := &Stack{arr: make([]*TreeNode, 0)}
	stack.Push(root)

	for !stack.IsEmpty() {
		node := stack.Pop()
		if node != nil {
			if node.Right != nil {
				stack.Push(node.Right)
			}

			stack.Push(node)
			stack.Push(nil)

			if node.Left != nil {
				stack.Push(node.Left)
			}
		} else {
			n := stack.Pop()
			n.Left = nil
			n.Right = nil
			tail.Right = n
			tail = n
		}
	}
	return dummy.Right
}
