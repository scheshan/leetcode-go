package m17_09

import "container/heap"

type Heap struct {
	arr []int
}

func (t *Heap) Len() int {
	return len(t.arr)
}

func (t *Heap) Less(i, j int) bool {
	return t.arr[i] < t.arr[j]
}

func (t *Heap) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *Heap) Push(x interface{}) {
	t.arr = append(t.arr, x.(int))
}

func (t *Heap) Pop() interface{} {
	n := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return n
}

func (t *Heap) Peek() int {
	return t.arr[0]
}

func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}

	m := make(map[int]bool)
	h := &Heap{arr: make([]int, 0)}

	addNumber(h, m, 3)
	addNumber(h, m, 5)
	addNumber(h, m, 7)

	for i := 2; i < k; i++ {
		num := heap.Pop(h).(int)
		addNumber(h, m, num*3)
		addNumber(h, m, num*5)
		addNumber(h, m, num*7)
	}

	return heap.Pop(h).(int)
}

func addNumber(h heap.Interface, m map[int]bool, n int) {
	if m[n] {
		return
	}
	m[n] = true
	heap.Push(h, n)
}
