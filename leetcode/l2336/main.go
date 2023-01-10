package main

import (
	"container/heap"
	"fmt"
)

type SmallestInfiniteSet struct {
	min int
	arr []int
	m   map[int]bool
}

func (t *SmallestInfiniteSet) Len() int {
	return len(t.arr)
}

func (t *SmallestInfiniteSet) Less(i, j int) bool {
	return t.arr[i] < t.arr[j]
}

func (t *SmallestInfiniteSet) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *SmallestInfiniteSet) Push(x interface{}) {
	t.arr = append(t.arr, x.(int))
}

func (t *SmallestInfiniteSet) Pop() interface{} {
	n := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return n
}

func Constructor() SmallestInfiniteSet {
	set := SmallestInfiniteSet{
		min: 1,
		arr: make([]int, 0),
		m:   make(map[int]bool),
	}
	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	res := 0

	if len(this.arr) > 0 {
		res = heap.Pop(this).(int)
		delete(this.m, res)
	} else {
		res = this.min
		this.min++
	}

	return res
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.min || this.m[num] {
		return
	}

	this.m[num] = true
	heap.Push(this, num)
}

func main() {
	s := Constructor()
	s1 := &s

	s1.AddBack(2)
	fmt.Println(s1.PopSmallest())
	fmt.Println(s1.PopSmallest())
	fmt.Println(s1.PopSmallest())
	s1.AddBack(1)
	fmt.Println(s1.PopSmallest())
	fmt.Println(s1.PopSmallest())
}
