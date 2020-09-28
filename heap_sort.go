package main

import "time"

var contador int
var iniheap int64
var finheap int64
var tempoheap float64

func heapSort(arr []int) []int {
	iniheap = time.Now().UnixNano()

	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		contador++
		siftDown(arr, i, len(arr))
		contador++
	}
	contador++
	for i := len(arr) - 1; i > 0; i-- {
		contador++
		// swap the root of the max-heap with the last item
		arr[0], arr[i] = arr[i], arr[0]
		contador++
		// fix heap
		siftDown(arr, 0, i)
		contador++
	}

	finheap = time.Now().UnixNano()

	tempoheap = float64(iniheap-finheap) / 1000000

	return arr
}

func siftDown(heap []int, lo, hi int) {
	root := lo
	contador++
	for {
		contador++
		child := root*2 + 1
		contador++
		if child >= hi {
			contador++
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			contador++
			child++
			contador++
		}
		if heap[root] < heap[child] {
			contador++
			heap[root], heap[child] = heap[child], heap[root]
			contador++
			root = child
			contador++
		} else {
			contador++
			break
		}

	}
}
