package p328

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func oddEvenList(head *ListNode) *ListNode {
	oddDummy := &ListNode{}
	oddTail := oddDummy
	evenDummy := &ListNode{}
	evenTail := evenDummy

	for head != nil {
		oddTail.Next = head
		oddTail = oddTail.Next
		head = head.Next
		if head != nil {
			evenTail.Next = head
			evenTail = evenTail.Next
			head = head.Next
		}
	}
	evenTail.Next = nil
	oddTail.Next = evenDummy.Next
	return oddDummy.Next
}
