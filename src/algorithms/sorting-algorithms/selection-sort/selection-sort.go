package main

type Type interface {
	int | int32 | int64 | string
}

func swap[T Type](firstPointer *T, secondPointer *T) {
	temp := *firstPointer
	*firstPointer = *secondPointer
	*secondPointer = temp
}

func selectionSort[T Type](array []T) {
	var n int = len(array)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		// swap(&array[min], &array[i])
		array[min], array[i] = array[i], array[min]
	}
}
