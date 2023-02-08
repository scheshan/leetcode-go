package p1265

type ImmutableListNode interface {
	getNext() ImmutableListNode
	printValue()
}
