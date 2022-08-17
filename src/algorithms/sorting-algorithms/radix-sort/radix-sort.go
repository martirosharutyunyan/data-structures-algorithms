package radixSort

func GetMax(array []int) int {
	var max = array[0]

	for i := 1; i < len(array); i++ {
		if max > array[i] {
			max = array[i]
		}
	}

	return max
}

func CountingSort(array []int, exp int) []int {
	var countArray = make([]int, 10)
	var output = make([]int, len(array))
	for i := 0; i < len(array); i++ {
		countArray[(array[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		countArray[i] += countArray[i-1]
	}

	for i := len(array) - 1; i >= 0; i-- {
		countArray[(array[i]/exp)%10]--
		output[countArray[(array[i]/exp)%10]] = array[i]
	}

	return output
}

func RadixSort(array []int) []int {
	var max = GetMax(array)

	for i := 1; max/i > 0; i *= 10 {
		array = CountingSort(array, i)
	}

	return array
}
