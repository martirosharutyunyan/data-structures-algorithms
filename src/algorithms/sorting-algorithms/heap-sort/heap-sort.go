package heapSort

func Heapify(arr []int, N, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < N && arr[l] > arr[largest] {
		largest = l
	}

	if r < N && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, N, largest)
	}
}

func HeapSort(arr []int, N int) {
	for i := N/2 - 1; i > -1; i-- {
		Heapify(arr, N, i)
	}

	for i := N - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		Heapify(arr, i, 0)
	}
}
