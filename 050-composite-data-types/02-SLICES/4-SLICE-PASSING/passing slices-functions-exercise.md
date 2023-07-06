Exercise: Passing Slices between Functions

Instructions:
1. Write a function called `doubleSlice` that takes a slice of integers as input and doubles the value of each element in the slice.
2. Write another function called `modifySlice` that takes a slice of integers as input and appends the value 5 to the slice.
3. In the `main` function:
   a. Create a slice of integers and populate it with some values of your choice.
   b. Call the `doubleSlice` function and pass the slice as an argument.
   c. Call the `modifySlice` function and pass the slice as an argument.
   d. Print the modified slice after both function calls.

Solution:

```go
package main

import "fmt"

func doubleSlice(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

func modifySlice(slice []int) {
	slice = append(slice, 5)
}

func main() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Call the doubleSlice function
	doubleSlice(numbers)

	// Call the modifySlice function
	modifySlice(numbers)

	// Print the modified slice
	fmt.Println(numbers)
}
```

Explanation:
In this exercise, we have three functions: `doubleSlice`, `modifySlice`, and `main`. The `doubleSlice` function takes a slice of integers as input and doubles the value of each element in the slice. The `modifySlice` function appends the value 5 to the given slice of integers. In the `main` function, we create a slice called `numbers` and populate it with some initial values. We then call the `doubleSlice` function, passing the `numbers` slice as an argument, which modifies the elements in-place by doubling their values. Next, we call the `modifySlice` function, passing the `numbers` slice as an argument. However, note that the `modifySlice` function doesn't modify the slice itself; it only appends the value 5 to a copy of the slice within the function. Finally, we print the modified `numbers` slice after both function calls.

Expected Output:
```
[2 4 6 8 10 5]
```

The output shows the modified `numbers` slice, where each element has been doubled by the `doubleSlice` function, and the value 5 has been appended by the `modifySlice` function.