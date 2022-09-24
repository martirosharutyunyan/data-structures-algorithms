package interpolationSearch

import (
	"fmt"
)

func InterpolationSearch(array []int, elem int) int {
	size := len(array)
	low := 0
	high := size - 1
	var middle int
	for (array[high] != array[low]) && (elem >= array[low]) && (elem <= array[high]) {
		middle = low + ((elem - array[low]) * (high - low) / (array[high] - array[low]))
		fmt.Println(middle, array[middle], elem, low, array[low], high, array[high])
		if array[middle] < elem {
			low = middle + 1
		} else if elem < array[middle] {
			high = middle - 1
		} else {
			return middle
		}
	}

	if elem == array[low] {
		return low
	}

	return -1
}
