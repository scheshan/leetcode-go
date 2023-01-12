package l919

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

type CBTInserter struct {
	arr []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	res := CBTInserter{
		arr: make([]*TreeNode, 0),
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			res.arr = append(res.arr, node)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}
	}

	return res
}

func (this *CBTInserter) Insert(v int) int {
	node := &TreeNode{Val: v}
	this.arr = append(this.arr, node)

	ind := (len(this.arr) - 2) >> 1
	parent := this.arr[ind]

	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
	}

	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.arr[0]
}
