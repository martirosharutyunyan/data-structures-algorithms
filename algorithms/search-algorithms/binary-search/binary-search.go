package main

import (
	"fmt"
)

func binarySearchRecursive(array []int, left int, right int, elem int) int {
	if right >= left {
		middle := left + (right-left)/2
		if array[middle] == elem {
			return middle
		}

		if array[middle] > elem {
			return binarySearchRecursive(array, left, middle-1, elem)
		}

		return binarySearchRecursive(array, middle+1, right, elem)
	}
	return -1
}

func binarySearchIterative(array []int, elem int) int {
	left := 0
	right := len(array) - 1
	var middle int

	for right >= left {
		middle = left + (right-left)/2

		if array[middle] == elem {
			return middle
		}

		if array[middle] > elem {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	elem := 7

	index := binarySearchRecursive(array, 0, len(array)-1, elem)
	index = binarySearchIterative(array, elem)
	fmt.Println(index)
}
