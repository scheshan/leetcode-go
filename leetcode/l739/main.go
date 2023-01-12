package l739

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

func dailyTemperatures(temperatures []int) []int {
	stack := &Stack{arr: make([]int, 0)}
	res := make([]int, len(temperatures))

	for i, n := range temperatures {
		for stack.Len() > 0 && temperatures[stack.Peek()] < n {
			pre := stack.Pop()
			res[pre] = i - pre
		}

		stack.Push(i)
	}

	return res
}
