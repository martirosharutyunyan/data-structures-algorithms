package stack

var stackArray []any

// stackPush push to first index of array
func stackPush(n any) {
	stackArray = append([]any{n}, stackArray...)
}

// stackLength return length of array
func stackLength() int {
	return len(stackArray)
}

// stackPeak return last input of array
func stackPeak() any {
	return stackArray[0]
}

// stackEmpty check array is empty or not
func stackEmpty() bool {
	return len(stackArray) == 0
}

// stackPop return last input and remove it in array
func stackPop() any {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}
