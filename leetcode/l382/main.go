package l382

import (
	"github.com/scheshan/leetcode/common"
	"math/rand"
)

type ListNode = common.ListNode

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	s := Solution{head: head}
	return s
}

func (this *Solution) GetRandom() int {
	num := 1
	head := this.head

	var res int

	for head != nil {
		if rand.Intn(num) == 0 {
			res = head.Val
		}
		num++
		head = head.Next
	}

	return res
}
