package main

import (
	"fmt"

	merge "github.com/woonmapao/random-merge-function"
)

func main() {
	collection1 := []int{9, 7, 5, 1}
	collection2 := []int{2, 4, 6, 8}
	collection3 := []int{0, 3, 10}

	fmt.Println(merge.Merge(collection1, collection2, collection3))
}
