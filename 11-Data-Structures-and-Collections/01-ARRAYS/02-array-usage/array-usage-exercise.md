Exercise: Array Manipulation

Instructions:
1. Declare an array named "ages" with a length of 5 and initialize it with the following values: 18, 25, 30, 42, 55.
2. Print the value at index 3 of the "ages" array.
3. Declare another array named "newAges" and assign the values of the "ages" array to it.
4. Change the value at index 2 of the "newAges" array to 35.
5. Print the values of both the "ages" and "newAges" arrays to verify that the change affected only the "newAges" array.
6. Declare an array named "names" with a length of 3 and initialize it with the names "Alice", "Bob", and "Charlie".
7. Declare an array named "nameLengths" with a length equal to the number of elements in the "names" array.
8. Use a loop to iterate over the "names" array and assign the length of each name to the corresponding index in the "nameLengths" array.
9. Print the values of the "nameLengths" array.

Example Output:
```
Value at index 3 of the "ages" array: 42
Values of the "ages" array: [18 25 30 42 55]
Values of the "newAges" array: [18 25 35 42 55]
Values of the "nameLengths" array: [5 3 7]
```

Solution:
```go
package main

import "fmt"

func main() {
	// 1. Declare and initialize the "ages" array
	var ages = [5]int{18, 25, 30, 42, 55}

	// 2. Print the value at index 3 of the "ages" array
	fmt.Println("Value at index 3 of the \"ages\" array:", ages[3])

	// 3. Declare the "newAges" array and assign the values of the "ages" array
	var newAges = ages

	// 4. Change the value at index 2 of the "newAges" array
	newAges[2] = 35

	// 5. Print the values of both arrays
	fmt.Println("Values of the \"ages\" array:", ages)
	fmt.Println("Values of the \"newAges\" array:", newAges)

	// 6. Declare and initialize the "names" array
	var names = [3]string{"Alice", "Bob", "Charlie"}

	// 7. Declare the "nameLengths" array
	var nameLengths [3]int

	// 8. Assign the length of each name to the corresponding index in the "nameLengths" array
	for i, name := range names {
		nameLengths[i] = len(name)
	}

	// 9. Print the values of the "nameLengths" array
	fmt.Println("Values of the \"nameLengths\" array:", nameLengths)
}
```

Note: Encourage students to modify and experiment with the code to deepen their understanding of arrays in Go.