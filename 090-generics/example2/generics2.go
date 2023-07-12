package main

import "fmt"

// Comparable interface for type constraint
type Comparable interface {
	LessThan(other interface{}) bool
	Equals(other interface{}) bool
}

type MyInt int
type MyString string

// BinarySearch performs a generic binary search on a sorted slice
func BinarySearch[T Comparable](arr []T, target T) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid].Equals(target) {
			return mid
		} else if arr[mid].LessThan(target) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Element not found
}

// Implement the LessThan method for type int
func (i MyInt) LessThan(other interface{}) bool {
	if val, ok := other.(MyInt); ok {
		return i < val
	}
	return false
}

// Implement the Equals method for type int
func (i MyInt) Equals(other interface{}) bool {
	if val, ok := other.(MyInt); ok {
		return i == val
	}
	return false
}

// Implement the LessThan method for type string
func (s MyString) LessThan(other interface{}) bool {
	if val, ok := other.(MyString); ok {
		return s < val
	}
	return false
}

// Implement the Equals method for type string
func (s MyString) Equals(other interface{}) bool {
	if val, ok := other.(MyString); ok {
		return s == val
	}
	return false
}

func main() {
	// Binary search on a sorted integer slice
	intSlice := []MyInt{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var targetInt MyInt = 14
	index := BinarySearch(intSlice, targetInt)
	fmt.Printf("Index of %d in intSlice: %d\n", targetInt, index)

	// Binary search on a sorted string slice
	stringSlice := []MyString{"apple", "banana", "cherry", "date", "grape", "kiwi", "orange", "pineapple"}
	var targetString MyString = "kiwi"
	index = BinarySearch(stringSlice, targetString)
	fmt.Printf("Index of %s in stringSlice: %d\n", targetString, index)
}
