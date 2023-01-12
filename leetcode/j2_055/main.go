package j2_055

import (
	"github.com/scheshan/leetcode/common"
)

type TreeNode = common.TreeNode

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	res := BSTIterator{
		stack: make([]*TreeNode, 0),
	}
	res.push(root)
	return res
}

func (this *BSTIterator) Next() int {
	for {
		node := this.pop()
		if node == nil {
			return this.pop().Val
		}

		if node.Right != nil {
			this.push(node.Right)
		}

		this.push(node)
		this.push(nil)

		if node.Left != nil {
			this.push(node.Left)
		}
	}
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func (this *BSTIterator) pop() *TreeNode {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	return node
}

func (this *BSTIterator) push(node *TreeNode) {
	this.stack = append(this.stack, node)
}
