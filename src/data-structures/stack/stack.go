package stack

var stackArray []interface{}

// stackPush push to first index of array
func stackPush(n interface{}) {
	stackArray = append([]interface{}{n}, stackArray...)
}

// stackLength return length of array
func stackLength() int {
	return len(stackArray)
}

// stackPeak return last input of array
func stackPeak() interface{} {
	return stackArray[0]
}

// stackEmpty check array is empty or not
func stackEmpty() bool {
	return len(stackArray) == 0
}

// stackPop return last input and remove it in array
func stackPop() interface{} {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}
