package l328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	evenHead, oddHead := new(ListNode), new(ListNode)
	evenTail, oddTail := evenHead, oddHead

	i := 1
	for head != nil {
		if i&1 == 1 {
			evenTail.Next = head
			evenTail = evenTail.Next
		} else {
			oddTail.Next = head
			oddTail = oddTail.Next
		}

		head = head.Next
		i++
	}

	evenTail.Next = oddHead.Next
	oddTail.Next = nil

	return evenHead.Next
}
