package l703

import "container/heap"

type KthLargest struct {
	arr []int
	k   int
}

func (t *KthLargest) Len() int {
	return len(t.arr)
}

func (t *KthLargest) Less(i, j int) bool {
	return t.arr[i] < t.arr[j]
}

func (t *KthLargest) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *KthLargest) Push(x interface{}) {
	t.arr = append(t.arr, x.(int))
}

func (t *KthLargest) Pop() interface{} {
	n := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return n
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		arr: make([]int, 0, k),
		k:   k,
	}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.Len() < this.k {
		heap.Push(this, val)
	} else {
		if val > this.arr[0] {
			heap.Pop(this)
			heap.Push(this, val)
		}
	}
	return this.arr[0]
}
