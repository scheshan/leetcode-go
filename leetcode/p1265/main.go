package p1265

func printLinkedListInReverse(head ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	head.printValue()
}
