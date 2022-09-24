package jumpSearch

import (
	"math"
)

func JumpSearch(array []int, elem int) int {
	step := int(math.Sqrt(float64(elem)))

	prev := 0
	for array[int(math.Min(float64(step), float64(len(array))))-1] < elem {
		prev = step
		step += int(math.Sqrt(float64(elem)))

		if prev >= len(array) {
			return -1
		}
	}

	for array[prev] < elem {
		prev++

		if prev == int(math.Min(float64(step), float64(len(array)))) {
			return -1
		}
	}

	if array[prev] == elem {
		return prev
	}

	return -1
}
