package j2_041

type MovingAverage struct {
	arr   []int
	size  int
	total float64
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	ma := MovingAverage{
		size: size,
		arr:  make([]int, 0),
	}

	return ma
}

func (this *MovingAverage) Next(val int) float64 {
	this.arr = append(this.arr, val)
	this.total += float64(val)

	if len(this.arr) > this.size {
		this.total -= float64(this.arr[0])
		this.arr = this.arr[1:]
	}

	return this.total / float64(len(this.arr))
}
