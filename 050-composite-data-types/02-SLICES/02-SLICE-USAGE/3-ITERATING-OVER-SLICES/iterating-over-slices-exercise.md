Exercise: Iterating Over Slices

Instructions:
1. Create a new Go file named "slice_iteration.go" on your local machine.
2. Implement the following exercises in the file, following the instructions provided for each.
3. Test your code to ensure it produces the expected output.
4. Once you have completed the exercises, refer to the solution and explanation provided below.

Exercises:
1. Iterate over a slice of integers and print each element on a new line using the `for range` loop.

2. Iterate over a slice of strings and print each element along with its index using the `for range` loop. Format the output as "Index: Value".

3. Iterate over a slice of floats using a traditional `for` loop and calculate the sum of all elements. Print the sum at the end.

Solution:

```go
package main

import "fmt"

func main() {
	// Exercise 1: Iterating over a slice of integers
	intSlice := []int{10, 20, 30, 40, 50}
	for _, value := range intSlice {
		fmt.Println(value)
	}

	fmt.Println("--------------------")

	// Exercise 2: Iterating over a slice of strings with indices
	strSlice := []string{"apple", "banana", "cherry", "date"}
	for index, value := range strSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	fmt.Println("--------------------")

	// Exercise 3: Iterating over a slice of floats and calculating the sum
	floatSlice := []float64{3.14, 1.618, 2.718, 0.577}
	sum := 0.0
	for i := 0; i < len(floatSlice); i++ {
		sum += floatSlice[i]
	}
	fmt.Println("Sum:", sum)
}
```

Explanation:

- Exercise 1:
  - We define a slice of integers `intSlice` with some values.
  - Using the `for range` loop, we iterate over each element of the slice.
  - The index is ignored using the blank identifier `_`, and the value is printed.

- Exercise 2:
  - We define a slice of strings `strSlice` with some values.
  - Using the `for range` loop, we iterate over each element of the slice.
  - Both the index and the value are accessed and printed using `fmt.Printf`.

- Exercise 3:
  - We define a slice of floats `floatSlice` with some values.
  - Using a traditional `for` loop, we iterate over the slice using the index.
  - The value at each index is added to the `sum` variable.
  - Finally, we print the calculated sum.

Note: The solution provided assumes the code is written in a single file for simplicity. In a real-world scenario, it is recommended to split the code into separate files, adhere to Go package conventions, and handle error checking as needed.