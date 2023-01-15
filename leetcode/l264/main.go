package l264

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

func nthUglyNumber(n int) int {
	h := &Heap{arr: make([]int, 0)}
	heap.Push(h, 1)

	visit := make(map[int]bool)

	res := 0
	for i := 0; i < n; i++ {
		res = heap.Pop(h).(int)
		addToHeap(h, res*2, visit)
		addToHeap(h, res*3, visit)
		addToHeap(h, res*5, visit)
	}

	return res
}

func addToHeap(h heap.Interface, num int, visit map[int]bool) {
	if visit[num] {
		return
	}
	visit[num] = true
	heap.Push(h, num)
}
