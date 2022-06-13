package linearSearch

type Type interface {
	int | int32 | int64 | string
}

func linearSearch[T Type](array []T, elem T) int {
	for i := 0; i < len(array); i++ {
		if array[i] == elem {
			return i
		}
	}

	return -1
}
