package main

type Type interface {
	int | int32 | int64 | string
}

func mergeSort[T Type](array T[]) T[]{
	
}


func main() {
	var array []int = []int{64, 25, 12, 22, 11}
	mergeSort(array)
	fmt.Println(array)
}
