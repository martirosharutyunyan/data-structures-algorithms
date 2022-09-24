package linearSearch

import "github.com/martirosharutyunyan/data-structures-algorithms/src/utils"

func LinearSearch[T utils.Type](array []T, elem T) int {
	for i := 0; i < len(array); i++ {
		if array[i] == elem {
			return i
		}
	}

	return -1
}
