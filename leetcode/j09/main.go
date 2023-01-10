package j09

type Stack []int

func (t *Stack) Push(v int) {
	arr := *t
	*t = append(arr, v)
}

func (t *Stack) Pop() int {
	arr := *t
	n := arr[len(arr)-1]
	*t = arr[:len(arr)-1]
	return n
}

func (t *Stack) Len() int {
	return len(*t)
}

type CQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() CQueue {
	s1 := Stack(make([]int, 0))
	s2 := Stack(make([]int, 0))

	q := CQueue{
		s1: &s1,
		s2: &s2,
	}
	return q
}

func (this *CQueue) AppendTail(value int) {
	this.s2.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s1.Len() == 0 {
		for this.s2.Len() != 0 {
			this.s1.Push(this.s2.Pop())
		}
	}

	if this.s1.Len() > 0 {
		return this.s1.Pop()
	}

	return -1
}
