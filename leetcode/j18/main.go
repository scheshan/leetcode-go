package j18

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	tail := dummy

	for head != nil {
		if head.Val != val {
			tail.Next = head
			tail = tail.Next
		}

		head = head.Next
	}

	tail.Next = nil
	return dummy.Next
}
