package timSort

import (
	"github.com/martirosharutyunyan/data-structures-algorithms/src/utils"
	"math"
)

var min int = 16

func insertionSort[T utils.Type](array []T, left int, right int) {
	var j int
	var key T
	for i := left + 1; i <= right; i++ {
		key = array[i]
		j = i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}

func merge[T utils.Type](array []T, left int, middle int, right int) {
	n1 := middle - left + 1
	n2 := right - middle

	l := make([]T, n1)
	r := make([]T, n2)

	for i := 0; i < n1; i++ {
		l[i] = array[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = array[middle+i+1]
	}

	i, j, k := 0, 0, left

	for i < n1 && j < n2 {
		if l[i] <= r[j] {
			array[k] = l[i]
			i++
		} else {
			array[k] = r[j]
			j++
		}
		k++
	}

	for i < len(l) {
		array[k] = l[i]
		i++
		k++
	}

	for j < len(r) {
		array[k] = r[j]
		j++
		k++
	}

}

func TimSort[T utils.Type](array []T) {
	n := len(array)
	minMerge := 32
	for i := 0; i < n; i += min {
		insertionSort(array, i, int(math.Min(float64(i+minMerge-1), float64(n-1))))
	}

	for size := min; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			middle := left + size - 1
			right := int(math.Min(float64(left+(2*size)-1), float64(n-1)))

			if middle < right {
				merge(array, left, middle, right)
			}
		}
	}
}
