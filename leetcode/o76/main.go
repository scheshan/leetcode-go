package o76

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

func findKthLargest(nums []int, k int) int {
	h := &Heap{
		arr: make([]int, 0, k),
	}

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if num > h.Peek() {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}

	return h.Peek()
}
