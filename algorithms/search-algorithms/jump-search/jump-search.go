package main

import (
	"fmt"
	"math"
)

func jumpSearch(array []int, elem int) int {
	step := int(math.Sqrt(float64(elem)))

	prev := 0
	for array[int(math.Min(float64(step), float64(len(array))))-1] < elem {
		prev = step
		step += int(math.Sqrt(float64(elem)))

		if prev >= len(array) {
			return -1
		}
	}

	for array[prev] < elem {
		prev++

		if prev == int(math.Min(float64(step), float64(len(array)))) {
			return -1
		}
	}

	if array[prev] == elem {
		return prev
	}

	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	elem := 7

	index := jumpSearch(array, elem)
	fmt.Println(index)
}
