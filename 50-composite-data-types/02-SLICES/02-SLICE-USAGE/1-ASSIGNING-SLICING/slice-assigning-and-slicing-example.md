Certainly! Here's a Go program that exemplifies the concepts we covered in the lesson on working with slices:

```go
package main

import "fmt"

func main() {
	// Declaring an array using an array literal
	arr := [5]int{1, 2, 3, 4, 5}

	// Taking the slice of a slice
	slice := arr[2:4]

	// How length and capacity are calculated
	length := len(slice)
	capacity := cap(slice)

	fmt.Println("Original Array:", arr)
	fmt.Println("Original Slice:", slice)
	fmt.Println("Length:", length)
	fmt.Println("Capacity:", capacity)

	// Calculating the new length and capacity
	newSlice := slice[1:3]
	newLength := len(newSlice)
	newCapacity := cap(newSlice)

	fmt.Println("New Slice:", newSlice)
	fmt.Println("New Length:", newLength)
	fmt.Println("New Capacity:", newCapacity)

	// Potential consequence of making changes to a slice
	slice[0] = 10

	fmt.Println("Modified Slice:", slice)
	fmt.Println("Original Array (after modification):", arr)

	// Runtime error showing index out of range
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Runtime Error:", r)
		}
	}()
	element := slice[10]
	fmt.Println("Accessing index 10 of the slice:", element)
}
```

Explanation:
1. The program starts by declaring an array named `arr` with integers using an array literal.
2. It then takes a slice of `arr` from index 2 to 4 (excluding 4) and assigns it to the variable `slice`.
3. The program demonstrates how to calculate the length and capacity of the slice using the `len()` and `cap()` functions.
4. Next, it creates a new slice `newSlice` from `slice` by taking a range of elements from index 1 to 3 (excluding 3).
5. The program showcases the calculation of the new length and capacity of the `newSlice`.
6. It modifies an element of the `slice` to demonstrate the potential consequence of making changes to a slice.
7. The program then prints the modified `slice` and the original `arr` to observe the impact of the modification.
8. Lastly, the program attempts to access an index out of range in the `slice`, resulting in a runtime error. It handles the error using a deferred function.

When you run this program, you should see the following output:

```
Original Array: [1 2 3 4 5]
Original Slice: [3 4]
Length: 2
Capacity: 3
New Slice: [4]
New Length: 1
New Capacity: 2
Modified Slice: [10 4]
Original Array (after modification): [1 2 10 4 5]
Runtime Error: runtime error: index out of range [10] with length 2
```

The output demonstrates the various concepts we covered, including declaring an array, taking slices, calculating length and capacity, modifying slices, and handling runtime errors when accessing indices out of range.