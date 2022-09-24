package quickSort

import "github.com/martirosharutyunyan/data-structures-algorithms/src/utils"

func QuickSort[T utils.Type](array []T) []T {
	if len(array) < 2 {
		return array
	}

	var leftItems []T
	var rightItems []T
	var pivot = array[0]

	for i := 1; i < len(array); i++ {
		if array[i] < pivot {
			leftItems = append(leftItems, array[i])
		} else {
			rightItems = append(rightItems, array[i])
		}
	}

	var sortedLeftItems = QuickSort(leftItems)
	var sortedRightItems = QuickSort(rightItems)

	var leftItemsWithPivot = append(sortedLeftItems, pivot)
	var sortedArray = append(leftItemsWithPivot, sortedRightItems...)

	return sortedArray
}
