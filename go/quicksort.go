// Code taken from https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
package main

import "math/rand"

// reviewpad-an: critical
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	low, high := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[high] = a[high], a[pivot]

	for i := range a {
		if a[i] < a[high] {
			a[low], a[i] = a[i], a[low]
			low++
		}
	}

	a[low], a[high] = a[high], a[low]

	quicksort(a[:low])
	quicksort(a[low+1:])

	return a
}
