package exponentialSearch

import (
	"math"
)

func ExponentialSearch(array []int, elem int) int {
	if array[0] == elem {
		return 0
	}

	i := 1
	for i < len(array) && array[i] <= elem {
		i *= 2
	}

	return BinarySearch(array, i/2, int(math.Min(float64(i), float64(len(array)-1))), elem)
}

func BinarySearch(array []int, left int, right int, elem int) int {
	var middle int

	for right >= left {
		middle = left + (right-left)/2

		if array[middle] == elem {
			return middle
		}

		if array[middle] > elem {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
