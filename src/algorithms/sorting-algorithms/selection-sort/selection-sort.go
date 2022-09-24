package main

import "github.com/martirosharutyunyan/data-structures-algorithms/src/utils"

func swap[T utils.Type](firstPointer *T, secondPointer *T) {
	temp := *firstPointer
	*firstPointer = *secondPointer
	*secondPointer = temp
}

func SelectionSort[T utils.Type](array []T) {
	var n int = len(array)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		// swap(&array[min], &array[i])
		array[min], array[i] = array[i], array[min]
	}
}
