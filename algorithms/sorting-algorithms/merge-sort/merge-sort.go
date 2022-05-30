package main

import (
	"fmt"
	"math"
)

type Type interface {
	int | int32 | int64 | string
}

func merge[T Type](left []T, right []T) []T {
	var array = make([]T, len(left)+len(right))
	var i = 0
	var j = 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			array[i+j] = left[i]
			i++
		} else {
			array[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		array[i+j] = left[i]
		i++

	}

	for j < len(right) {
		array[i+j] = right[j]
		j++
	}

	return array
}

func mergeSort[T Type](array []T) []T {
	if len(array) < 2 {
		return array
	}

	var middle = len(array) / 2
	var leftItems = mergeSort(array[:middle])
	var rightItems = mergeSort(array[middle:])
	return merge(leftItems, rightItems)
}

func MergeIter[T Type](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:int(math.Min(float64(i+2*step), float64(len(items))))])
			copy(items[i:], tmp)
		}
	}
	return items
}

func main() {
	var array []int = []int{64, 25, 12, 22, 11}
	array = mergeSort(array)
	fmt.Println(array)
}
