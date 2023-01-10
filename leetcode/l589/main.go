package l589

import "github.com/scheshan/leetcode/common"

type Node = common.Node

type Stack struct {
	arr []*Node
}

func (t *Stack) Push(node *Node) {
	t.arr = append(t.arr, node)
}

func (t *Stack) Pop() *Node {
	node := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return node
}

func (t *Stack) IsEmpty() bool {
	return len(t.arr) == 0
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := &Stack{arr: make([]*Node, 0)}
	stack.Push(root)

	res := make([]int, 0)

	for !stack.IsEmpty() {
		node := stack.Pop()
		if node != nil {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack.Push(node.Children[i])
			}

			stack.Push(node)
			stack.Push(nil)
		} else {
			node = stack.Pop()
			res = append(res, node.Val)
		}
	}

	return res
}
