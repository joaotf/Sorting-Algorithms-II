package main

import (
	"math/rand"
	"time"
)

func quicksort(a []int) ([]int, int, float64) {

	var inicio int64
	var final int64
	var tempofinal float64

	inicio = time.Now().UnixNano()
	var x int

	if len(a) < 2 {
		x++
		final = time.Now().UnixNano()
		tempofinal = float64(final-inicio) / 1000000
		return a, x, tempofinal
	}
	x++

	left, right := 0, len(a)-1
	x++

	pivot := rand.Int() % len(a)
	x++

	a[pivot], a[right] = a[right], a[pivot]
	x++

	for i, _ := range a {
		x++
		if a[i] < a[right] {
			x++
			a[left], a[i] = a[i], a[left]
			x++
			left++
			x++
		}
	}

	a[left], a[right] = a[right], a[left]
	x++
	quicksort(a[:left])
	x++
	quicksort(a[left+1:])
	x++

	final = time.Now().UnixNano()

	tempofinal = float64(final-inicio) / 1000000
	return a, x, tempofinal
}
