package main

import (
	"fmt"
)

func buckerSort(array []float64) {
	var buckets = make([][]float64, len(array))

	for i := 0; i < len(array); i++ {
		buckets[int(array[i]*float64(len(array)))] = append(buckets[int(array[i]*float64(len(array)))], array[i])
	}

	for i := 0; i < len(array); i++ {
		insertionSort(buckets[int(array[i]*float64(len(array)))])
	}

	index := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			array[index] = buckets[i][j]
			index++
		}
	}
}

func insertionSort(array []float64) {
	for i := 1; i < len(array); i++ {
		j := i - 1
		key := array[i]

		for j >= 0 && array[j] >= key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}

func main() {
	var array = []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	buckerSort(array)
	fmt.Println(array)
}
