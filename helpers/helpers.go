package helpers

import (
	"sort"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func binSearch(arr []int, t, low, high int) int {

	mid := (high + low) / 2

	if arr[mid] == t {

		return mid

	}

	if low >= high {

		return -1

	}

	if arr[mid] > t {

		return binSearch(arr, t, low, mid)

	}

	return binSearch(arr, t, low+1, high)

}

func BinarySearch(inputArray []int, finding int) int {
	sort.Ints(inputArray)
	return binSearch(inputArray, finding, 0, len(inputArray)-1)
}
