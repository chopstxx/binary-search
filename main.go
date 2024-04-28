package main

import (
	"fmt"

	"example.com/binary-search/binarySearch"
)

func main() {
	// Usage example
	slice := []int{1, 3, 5, 7, 9}
	target := 5

	index := binarySearch.FindIndex(slice, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the slice\n", target)
	}
}