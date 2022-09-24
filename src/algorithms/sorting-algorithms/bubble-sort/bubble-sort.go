package bubbleSort

import "github.com/martirosharutyunyan/data-structures-algorithms/src/utils"

func swap[T utils.Type](firstPointer *T, secondPointer *T) {
	temp := *firstPointer
	*firstPointer = *secondPointer
	*secondPointer = temp
}

func BubbleSort[T utils.Type](array []T) {
	var n int = len(array)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				swap(&array[j], &array[j+1])
			}
		}
	}
}
