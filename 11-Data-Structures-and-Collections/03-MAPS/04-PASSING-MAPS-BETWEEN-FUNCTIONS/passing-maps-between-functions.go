package main

import "fmt"

// Function to add key-value pairs to the map by value
func addToMapByValue(m map[int]string) {
	m[1] = "One"
	m[2] = "Two"
}

// Function to modify the map by reference
func modifyMapByReference(m map[int]string) {
	m[2] = "New Value"
}

func main() {
	// Declare a map
	myMap := make(map[int]string)

	// Call addToMapByValue and pass the map by value
	addToMapByValue(myMap)

	// Print the map after adding values
	fmt.Println("Map after adding values:", myMap)

	// Call modifyMapByReference and pass the map by reference
	modifyMapByReference(myMap)

	// Print the map after modification
	fmt.Println("Map after modification:", myMap)
}