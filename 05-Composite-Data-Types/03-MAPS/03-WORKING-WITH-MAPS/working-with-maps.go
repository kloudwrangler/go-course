package main

import (
	"fmt"
)

func main() {
	// Assigning values to a map
	myMap := make(map[string]int)
	myMap["apple"] = 3
	myMap["banana"] = 5

	// Runtime error with nil maps
	// var nilMap map[string]int
	// Uncomment the line below to see the runtime error
	// nilMap["apple"] = 3

	// Retrieving values and testing existence
	quantity := myMap["apple"]
	if _, ok := myMap["banana"]; ok {
		fmt.Println("The key 'banana' exists in the map.")
	} else {
		fmt.Println("The key 'banana' does not exist in the map.")
	}

	// Retrieving values and testing the value for existence
	quantity, exists := myMap["apple"]
	if exists {
		fmt.Println("The quantity of 'apple' is", quantity)
	} else {
		fmt.Println("'apple' does not exist in the map")
	}

	// Iterating over a map using for range
	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}

	// Removing an item from a map
	delete(myMap, "apple")

	// Printing the updated map
	fmt.Println("After removing 'apple':", myMap)
}
