# Exercise: Growing Slices Practice

1. Create a Go program that demonstrates the concepts of growing slices using the `append` function. Follow the instructions below:

   a. Declare an empty slice of integers called `numbers`.

   b. Prompt the user to enter a series of integers (separated by spaces) until they enter a non-integer value (e.g., "done" or any non-numeric input).

   c. Parse the user input and add each integer to the `numbers` slice using the `append` function.

   d. Print the contents of the `numbers` slice after each addition.

   e. Once the user enters a non-integer value, print the final `numbers` slice.

   f. Run the program and test it with different inputs to observe how the slice grows dynamically.

Solution:

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []int

	fmt.Println("Enter a series of integers (separated by spaces). Enter a non-integer value to stop:")

	var input string
	for {
		fmt.Scan(&input)

		if input == "done" {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			break
		}

		numbers = append(numbers, number)
		fmt.Println("Current numbers:", numbers)
	}

	fmt.Println("Final numbers:", numbers)
}
```

Explanation:

1. We declare an empty slice of integers called `numbers` to store the user input.

2. We use a loop to continuously read user input until they enter a non-integer value or the string "done".

3. Within the loop, we convert the input string to an integer using `strconv.Atoi()`. If the conversion is successful, we append the integer to the `numbers` slice using `append(numbers, number)`.

4. After each addition, we print the current contents of the `numbers` slice using `fmt.Println("Current numbers:", numbers)`.

5. Once the user enters a non-integer value or "done", we break out of the loop and print the final `numbers` slice using `fmt.Println("Final numbers:", numbers)`.

6. Running the program allows the user to enter a series of integers, and they can observe how the slice dynamically grows as each integer is added.

Note: Remind students to import the necessary packages (`fmt` and `strconv`) and save the program with a `.go` extension before running it.