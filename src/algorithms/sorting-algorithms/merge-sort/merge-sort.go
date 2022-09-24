package mergeSort

import (
	"github.com/martirosharutyunyan/data-structures-algorithms/src/utils"
	"math"
)

func Merge[T utils.Type](left []T, right []T) []T {
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

func MergeSort[T utils.Type](array []T) []T {
	if len(array) < 2 {
		return array
	}

	var middle = len(array) / 2
	var leftItems = MergeSort(array[:middle])
	var rightItems = MergeSort(array[middle:])
	return Merge(leftItems, rightItems)
}

func MergeIterative[T utils.Type](items []T) []T {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			tmp := Merge(items[i:i+step], items[i+step:int(math.Min(float64(i+2*step), float64(len(items))))])
			copy(items[i:], tmp)
		}
	}
	return items
}

// second variant
// func Merge[T Type](array []T, left int, middle int, right int) {
// 	n1 := middle - left + 1
// 	n2 := right - middle

// 	l := make([]T, n1)
// 	r := make([]T, n2)

// 	for i := 0; i < n1; i++ {
// 		l[i] = array[left+i]
// 	}
// 	for i := 0; i < n2; i++ {
// 		r[i] = array[middle+i+1]
// 	}

// 	i, j, k := 0, 0, left

// 	for i < n1 && j < n2 {
// 		if l[i] <= r[j] {
// 			array[k] = l[i]
// 			i++
// 		} else {
// 			array[k] = r[j]
// 			j++
// 		}
// 		k++
// 	}

// 	for i < len(l) {
// 		array[k] = l[i]
// 		i++
// 		k++
// 	}

// 	for j < len(r) {
// 		array[k] = r[j]
// 		j++
// 		k++
// 	}

// }

// func MergeSort[T Type](array []T, left int, right int) {
// 	if left >= right {
// 		return
// 	}
// 	middle := left + (right-left)/2
// 	MergeSort(array, left, middle)
// 	MergeSort(array, middle+1, right)
// 	Merge(array, left, middle, right)
// }
