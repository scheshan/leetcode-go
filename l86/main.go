package l86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	smallHead, bigHead := new(ListNode), new(ListNode)
	smallTail, bigTail := smallHead, bigHead

	for head != nil {
		if head.Val < x {
			smallTail.Next = head
			smallTail = smallTail.Next
		} else {
			bigTail.Next = head
			bigTail = bigTail.Next
		}

		head = head.Next
	}

	bigTail.Next = nil
	smallTail.Next = bigHead.Next
	return smallHead.Next
}
