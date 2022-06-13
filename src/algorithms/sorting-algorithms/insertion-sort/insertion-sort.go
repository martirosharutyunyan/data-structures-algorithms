package insertionSort

type Type interface {
	int | int32 | int64 | string
}

func insertionSort[T Type](array []T) {
	var j int
	var key T
	for i := 1; i < len(array); i++ {
		key = array[i]
		j = i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}