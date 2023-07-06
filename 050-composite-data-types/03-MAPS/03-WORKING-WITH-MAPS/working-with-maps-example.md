# program that exemplifies all the topics covered in the lesson on working with maps in Go:

```go
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
	var nilMap map[string]int
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
```

Explanation:
1. The program starts by creating a map named `myMap` and assigning values to it using the key-value syntax.
2. Then, a nil map named `nilMap` is declared, which demonstrates the runtime error that occurs when trying to assign a value to a nil map.
3. The program retrieves a value from `myMap` using the key "apple" and tests for the existence of the key "banana" using the comma-ok idiom.
4. Next, the program retrieves the value of "apple" and checks its existence using the two-variable syntax.
5. A for range loop is used to iterate over `myMap` and print each key-value pair.
6. The program deletes the key-value pair with the key "apple" using the `delete` function.
7. Finally, the updated map is printed to show that "apple" has been removed.

By running this program, participants can observe and understand the concepts of assigning values to a map, handling runtime errors with nil maps, retrieving values and testing existence, iterating over a map, and removing items from a map.

Feel free to modify the program or add more functionality to further illustrate the concepts and provide additional examples.