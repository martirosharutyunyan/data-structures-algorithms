package main

import "fmt"

type Type interface {
	int | int32 | int64 | string
}

func linearSearch[T Type](array []T, elem T) int {
	for i := 0; i < len(array); i++ {
		if array[i] == elem {
			return i
		}
	}

	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	elem := 7

	index := linearSearch(array, elem)

	fmt.Println(index)
}
