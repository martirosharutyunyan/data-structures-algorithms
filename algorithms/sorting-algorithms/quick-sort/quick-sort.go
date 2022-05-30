package main

import "fmt"

type Type interface {
	int | int32 | int64 | string
}

func quickSort[T Type](array []T) []T {
	if len(array) < 2 {
		return array
	}

	var leftItems = []T{}
	var rightItems = []T{}
	var pivot = array[0]

	for i := 1; i < len(array); i++ {
		if array[i] < pivot {
			leftItems = append(leftItems, array[i])
		} else {
			rightItems = append(rightItems, array[i])
		}
	}

	var sortedLeftItems = quickSort(leftItems)
	var sortedRightItems = quickSort(rightItems)

	var leftItemsWithPivot = append(sortedLeftItems, pivot)
	var sortedArray = append(leftItemsWithPivot, sortedRightItems...)

	return sortedArray
}

func main() {
	var array []int = []int{64, 25, 12, 22, 11}
	array = quickSort(array)
	fmt.Println(array)
}
