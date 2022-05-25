package main

import (
	"container/list"
)

type MovingAverage struct {
	queue *list.List
	size  int
	total float64
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	avg := MovingAverage{}
	avg.queue = list.New()
	avg.size = size
	return avg
}

func (this *MovingAverage) Next(val int) float64 {
	if this.queue.Len() >= this.size {
		e := this.queue.Front()
		this.queue.Remove(e)

		this.total = this.total - e.Value.(float64)
	}

	v := float64(val)
	this.total = this.total + v
	this.queue.PushBack(v)

	return this.total / float64(this.queue.Len())
}
