package queue

var ListQueue []any

// EnQueue it will be added new value into our list
func EnQueue(n any) {
	ListQueue = append(ListQueue, n)
}

// DeQueue it will be removed the first value that added into the list
func DeQueue() any {
	data := ListQueue[0]
	ListQueue = ListQueue[1:]
	return data
}

// FrontQueue return the Front value
func FrontQueue() any {
	return ListQueue[0]
}

// BackQueue return the Back value
func BackQueue() any {
	return ListQueue[len(ListQueue)-1]
}

// LenQueue will return the length of the queue list
func LenQueue() int {
	return len(ListQueue)
}

// IsEmptyQueue check our list is empty or not
func IsEmptyQueue() bool {
	return len(ListQueue) == 0
}
