package j2_078

import (
	"container/heap"
	"github.com/scheshan/leetcode/common"
)

type ListNode = common.ListNode

type Heap []*ListNode

func (t *Heap) Len() int {
	return len(*t)
}

func (t *Heap) Less(i, j int) bool {
	return (*t)[i].Val < (*t)[j].Val
}

func (t *Heap) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *Heap) Push(x interface{}) {
	*t = append(*t, x.(*ListNode))
}

func (t *Heap) Pop() interface{} {
	node := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	s := Heap(make([]*ListNode, 0, len(lists)))
	h := &s
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		tail.Next = node
		tail = node

		node = node.Next
		if node != nil {
			heap.Push(h, node)
		}
	}

	return dummy.Next
}
