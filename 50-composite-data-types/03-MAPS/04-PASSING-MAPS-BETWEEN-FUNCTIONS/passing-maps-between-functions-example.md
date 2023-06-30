Certainly! Here's a program that incorporates the topics covered in the lesson on passing maps between functions:

```go
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
```

Explanation:
1. The program defines two functions: `addToMapByValue` and `modifyMapByReference`.
2. `addToMapByValue` takes a map by value and adds key-value pairs to it.
3. `modifyMapByReference` takes a map by reference and modifies one of its values.
4. In the `main` function, a map named `myMap` is declared using the `make` function.
5. The `addToMapByValue` function is called, passing `myMap` by value.
6. The map is printed after adding values.
7. The `modifyMapByReference` function is called, passing `myMap` by reference.
8. The map is printed again after modification.

When you run the program, the output should be:

```
Map after adding values: map[1:One 2:Two]
Map after modification: map[1:One 2:New Value]
```

This program demonstrates passing maps between functions by value and reference. The `addToMapByValue` function modifies the map by value, ensuring that the original map remains unaffected. In contrast, the `modifyMapByReference` function modifies the map by reference, which directly modifies the original map.

Feel free to run the program and explore the behavior of maps in Go.