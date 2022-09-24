package utils

func GetPointer[T any](value T) *T {
	return &value
}

func GetValueByPointer[T any](value *T) T {
	return *value
}

type Type interface {
	int | int32 | int64 | float64 | string
}
