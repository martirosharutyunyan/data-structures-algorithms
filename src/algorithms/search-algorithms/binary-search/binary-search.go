package binarySearch

func BinarySearchRecursive(array []int, left int, right int, elem int) int {
	if right >= left {
		middle := left + (right-left)/2
		if array[middle] == elem {
			return middle
		}

		if array[middle] > elem {
			return BinarySearchRecursive(array, left, middle-1, elem)
		}

		return BinarySearchRecursive(array, middle+1, right, elem)
	}
	return -1
}

func BinarySearchIterative(array []int, elem int) int {
	left := 0
	right := len(array) - 1
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
