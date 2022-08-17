package bubbleSort

type Type interface {
	int | int32 | int64 | string
}

func swap[T Type](firstPointer *T, secondPointer *T) {
	temp := *firstPointer
	*firstPointer = *secondPointer
	*secondPointer = temp
}

func BubbleSort[T Type](array []T) {
	var n int = len(array)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				swap(&array[j], &array[j+1])
			}
		}
	}
}
