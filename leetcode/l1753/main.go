package l1753

import "container/heap"

type Heap struct {
	arr []int
}

func (t *Heap) Len() int {
	return len(t.arr)
}

func (t *Heap) Less(i, j int) bool {
	return t.arr[j] < t.arr[i]
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

func maximumScore(a int, b int, c int) int {
	h := &Heap{arr: []int{a, b, c}}
	heap.Init(h)

	res := 0
	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)

		x--
		y--

		if x > 0 {
			heap.Push(h, x)
		}
		if y > 0 {
			heap.Push(h, y)
		}
		res++
	}

	return res
}
