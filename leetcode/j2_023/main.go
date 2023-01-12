package j2_023

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	n1, n2 := headA, headB

	for n1 != n2 {
		if n1 == nil {
			n1 = headB
		} else {
			n1 = n1.Next
		}
		if n2 == nil {
			n2 = headA
		} else {
			n2 = n2.Next
		}
	}
	return n1
}
