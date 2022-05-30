package main

import "fmt"

var diapason = 1000

func getMax(array []int) int {
	var max = array[0]

	for i := 1; i < len(array); i++ {
		if max > array[i] {
			max = array[i]
		}
	}

	return max
}

func countingSort(array []int) {
	// var countArray = make([]int, diapason)
	// for i := 0; i < len(array); i++ {
	// 	countArray[array[i]]++
	// }

	// z := 0
	// for i, c := range countArray {
	// 	for c > 0 {
	// 		array[z] = i
	// 		c--
	// 		z++
	// 	}
	// }
}

func radixSort(array []int) {
	var max = getMax(array)

	for i := 1; max/i > 0; i *= 10 {
		// countingSort(array)
	}
}

func main() {
	var array []int = []int{64, 25, 12, 22, 11}
	radixSort(array)
	fmt.Println(array)
}
