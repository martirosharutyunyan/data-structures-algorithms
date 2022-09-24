package bucketSort

import insertionSort "github.com/martirosharutyunyan/data-structures-algorithms/src/algorithms/sorting-algorithms/insertion-sort"

func BucketSort(array []float64) {
	var buckets = make([][]float64, len(array))

	for i := 0; i < len(array); i++ {
		buckets[int(array[i]*float64(len(array)))] = append(buckets[int(array[i]*float64(len(array)))], array[i])
	}

	for i := 0; i < len(array); i++ {
		insertionSort.InsertionSort(buckets[int(array[i]*float64(len(array)))])
	}

	index := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			array[index] = buckets[i][j]
			index++
		}
	}
}
