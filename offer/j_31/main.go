package j_31

type Stack struct {
	arr []int
}

func (t *Stack) Push(v int) {
	t.arr = append(t.arr, v)
}

func (t *Stack) Pop() int {
	n := t.Peek()
	t.arr = t.arr[:len(t.arr)-1]
	return n
}

func (t *Stack) Peek() int {
	return t.arr[len(t.arr)-1]
}

func (t *Stack) Len() int {
	return len(t.arr)
}

func validateStackSequences(pushed []int, popped []int) bool {
	l := 0
	r := 0
	stack := &Stack{
		arr: make([]int, 0),
	}

	for r < len(popped) {
		if stack.Len() == 0 || stack.Peek() != popped[r] {
			if l >= len(pushed) {
				break
			}

			stack.Push(pushed[l])
			l++
		} else {
			stack.Pop()
			r++
		}
	}

	return r == len(popped)
}
