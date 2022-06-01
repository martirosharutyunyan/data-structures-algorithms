package main

import "fmt"

func buckerSort(array []float64) {
	var buckets = make([][]float64, len(array))

	fmt.Println(buckets)
}

func main() {
	var array = []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	buckerSort(array)
	fmt.Println(array)
}
