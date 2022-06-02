package main

import (
	"fmt"
)

var diapason = 10

func countingSort(array []int) []int {
	var countArray = make([]int, diapason)
	var output = make([]int, len(array))

	for i := 0; i < len(array); i++ {
		countArray[array[i]]++
	}

	for i := 1; i < diapason; i++ {
		countArray[i] += countArray[i-1]
	}

	for i := len(array) - 1; i >= 0; i-- {
		countArray[array[i]] -= 1
		output[countArray[array[i]]] = array[i]
	}

	return output
}

func main() {
	var array []int = []int{1, 4, 1, 2, 7, 5, 2}
	array = countingSort(array)
	fmt.Println(array)
}
