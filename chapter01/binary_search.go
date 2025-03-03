package chapter01

import (
	"constraints"
	"errors"
)

// BinarySearch binary search has a runtime of O(logn).
// Why? Because it half the search list every time.
// Logn has this property where it tells you how many times
// you have to divide k with n to get to 1.
// Log2(16)=4 -> 16 -> 8 -> 4 -> 2 -> 1
func BinarySearch[T constraints.Ordered](list []T, item T) (T, error) {
	var result T
	head := 0
	tail := len(list) - 1
	for head <= tail {
		middle := (head + tail)/2
		guess := list[middle]
		if guess == item {
			return guess, nil
		}
		if guess > item {
			tail = middle - 1
		} else {
			head = middle + 1
		}
	}
	return result, errors.New("not found")
}
