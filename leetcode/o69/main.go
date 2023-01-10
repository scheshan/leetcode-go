package main

func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr)

	for l < r {
		mid := l + (r-l)>>1
		if mid == 0 || mid == len(arr)-1 || (arr[mid-1] < arr[mid] && arr[mid+1] < arr[mid]) {
			return mid
		} else if arr[mid-1] > arr[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return -1
}
