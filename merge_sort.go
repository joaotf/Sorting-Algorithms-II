package main

import "time"

var count int
var inicio int64
var final int64
var tempofinal float64

func mergeSort(items []int) []int {

	var num = len(items)
	count++

	if num == 1 {
		count++
		final = time.Now().UnixNano()
		tempofinal = float64(final-inicio) / 1000000
		return items
	}

	middle := int(num / 2)
	count++

	var left = make([]int, middle)
	var right = make([]int, num-middle)

	count++
	count++
	for i := 0; i < num; i++ {
		count++
		if i < middle {
			count++
			left[i] = items[i]
			count++
		} else {
			count++
			right[i-middle] = items[i]
			count++
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) (result []int) {
	inicio = time.Now().UnixNano()

	result = make([]int, len(left)+len(right))
	count++

	i := 0
	count++
	for len(left) > 0 && len(right) > 0 {
		count++
		if left[0] < right[0] {
			count++
			result[i] = left[0]
			count++
			left = left[1:]
			count++
		} else {
			count++
			result[i] = right[0]
			count++
			right = right[1:]
			count++
		}
		i++
		count++
	}

	for j := 0; j < len(left); j++ {
		count++
		result[i] = left[j]
		count++
		i++
		count++
	}
	for j := 0; j < len(right); j++ {
		count++
		result[i] = right[j]
		count++
		i++
		count++
	}

	tempofinal = float64(final-inicio) / 1000000
	return
}
