package main

import (
	"fmt"

	"github.com/reviewpad/action-showcase/go/sort"
)

func main() {
	fmt.Println(sort.Quicksort([]int{5, 6, 7, 2, 1, 0}))
}
