package main

import (
	"fmt"
)

var diapason = 10

func countingSort(array []int) {
	var countArray = make([]int, diapason)
	for i := 0; i < len(array); i++ {
		countArray[array[i]]++
	}

	z := 0
	for i, c := range countArray {
		for c > 0 {
			array[z] = i
			c--
			z++
		}
	}
}

func main() {
	var array []int = []int{1, 4, 1, 2, 7, 5, 2}
	countingSort(array)
	fmt.Println(array)
}
