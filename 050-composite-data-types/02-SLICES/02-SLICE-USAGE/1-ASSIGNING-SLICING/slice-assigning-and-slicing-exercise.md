Exercise: Manipulating Slices

Instructions:
1. Create a Go program that demonstrates various operations on slices based on the concepts discussed in the lesson.
2. Follow the steps provided in the comments of the program to complete the exercise.
3. Compile and run the program to verify the results.

```go
package main

import "fmt"

func main() {
	// Step 1: Declare an array with integers from 1 to 10 using an array literal.
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Step 2: Take a slice of the array that includes elements from index 2 to 6 (excluding 6).
	slice := arr[2:6]

	// Step 3: Print the slice, its length, and its capacity.
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Step 4: Modify the element at index 1 of the slice to 20.
	slice[1] = 20

	// Step 5: Append a new element 30 to the slice.
	slice = append(slice, 30)

	// Step 6: Take a new slice of the existing slice from index 1 to 4 (excluding 4).
	newSlice := slice[1:4]

	// Step 7: Print the new slice, its length, and its capacity.
	fmt.Println("New Slice:", newSlice)
	fmt.Println("Length:", len(newSlice))
	fmt.Println("Capacity:", cap(newSlice))

	// Step 8: Modify the element at index 0 of the new slice to 40.
	newSlice[0] = 40

	// Step 9: Print the modified slice and the original slice.
	fmt.Println("Modified New Slice:", newSlice)
	fmt.Println("Original Slice:", slice)

	// Step 10: Access an element at index 10 of the original slice and handle any runtime error.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Runtime Error:", r)
		}
	}()
	fmt.Println("Accessing index 10 of the original slice:", slice[10])
}
```

Explanation:
1. In this exercise, we start by declaring an array named `arr` with integers from 1 to 10 using an array literal.
2. We then take a slice of `arr` named `slice`, which includes elements from index 2 to 6 (excluding 6).
3. The program prints the slice, its length, and its capacity to demonstrate how they are calculated.
4. Next, we modify the element at index 1 of `slice` to 20 and append a new element 30 to the slice.
5. We then take a new slice named `newSlice` from the existing `slice`, starting from index 1 and excluding index 4.
6. The program prints the new slice, its length, and its capacity.
7. We modify the element at index 0 of `newSlice` to 40 and print both the modified `newSlice` and the original `slice` to observe the potential consequence of making changes to a slice.
8. Finally, we attempt to access an element at index 10 of the original slice, which is out of range. We handle the runtime error using a deferred function.

Solution:
When you compile and run the program, you should see the following output:

```
Slice: [3 4 5 6]
Length: 4
Capacity: 8
New Slice: [4 5 6]
Length: 3


Capacity: 7
Modified New Slice: [40 5 6]
Original Slice: [3 20 30 40 5 6]
Runtime Error: runtime error: index out of range [10] with length 6
```

The output demonstrates the manipulation of slices, the calculation of their length and capacity, the consequence of modifying a slice, and the runtime error when accessing an index out of range.