Exercise: Pointers in Go

In this exercise, you will practice working with pointers in Go. Complete the following tasks to reinforce your understanding of pointers and their usage.

1. Declare and Initialize Pointers:
- Declare a pointer variable called `ptr` of type `*int`.
- Declare an integer variable called `num` and assign it a value of your choice.
- Initialize `ptr` with the address of `num` using the `&` operator.

2. Dereference and Modify Pointers:
- Print the value stored at the memory address pointed to by `ptr` using dereferencing.
- Modify the value of `num` by assigning a new value to the memory address pointed to by `ptr`.

3. Pass Pointers to Functions:
- Write a function called `increment` that takes a pointer to an integer as a parameter.
- Inside the `increment` function, increment the value of the integer pointed to by the parameter by 1.
- Call the `increment` function, passing the address of `num` as an argument.
- Print the value of `num` after calling the `increment` function.

4. Reflection:
- In a comment, explain why we pass the address of `num` to the `increment` function instead of directly passing `num` itself.

Note: Make sure to run your code and verify that it produces the expected output.

```go
package main

import "fmt"

func increment(ptr *int) {
	// Increment the value pointed to by ptr by 1
	*ptr = *ptr + 1
}

func main() {
	// Task 1: Declare and Initialize Pointers
	var ptr *int
	num := 42
	ptr = &num

	// Task 2: Dereference and Modify Pointers
	fmt.Println("Value of num:", *ptr)
	*ptr = 99

	// Task 3: Pass Pointers to Functions
	increment(ptr)
	fmt.Println("Value of num after increment:", num)

	// Task 4: Reflection
	// We pass the address of num to the increment function because we want to modify the actual value of num,
	// and not a copy of it. By passing the address (pointer), the function can directly modify the value at that memory location.
}
```

Once you have completed the exercise, run the code and verify that the output matches your expectations. Experiment further with pointers to gain more familiarity with their behavior and usage in Go.