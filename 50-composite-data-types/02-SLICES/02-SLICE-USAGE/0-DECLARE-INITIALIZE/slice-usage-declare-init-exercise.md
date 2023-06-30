Exercise:

Write a program that accepts a list of numbers from the user and calculates the average of those numbers using slices.

Solution:

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a list of numbers (space-separated): ")
	fmt.Scanln(&input)

	numbers := parseNumbers(input)
	average := calculateAverage(numbers)

	fmt.Printf("Average: %.2f\n", average)
}

func parseNumbers(input string) []float64 {
	strNumbers := strings.Split(input, " ")
	numbers := make([]float64, 0, len(strNumbers))

	for _, str := range strNumbers {
		num, err := strconv.ParseFloat(str, 64)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}
```

Explanation:

1. The program prompts the user to enter a list of numbers separated by spaces.
2. The `parseNumbers` function takes the user input as a string and splits it into individual number strings using `strings.Split()`.
3. It then iterates through the number strings, converts each string to a float64 using `strconv.ParseFloat()`, and appends it to the `numbers` slice.
4. The `calculateAverage` function calculates the sum of all numbers in the `numbers` slice by iterating through it and adding each element to the `sum` variable.
5. Finally, the average is calculated by dividing the sum by the length of the `numbers` slice, and the result is printed to the console with two decimal places using `fmt.Printf()`.

Example Run:

```
Enter a list of numbers (space-separated): 5.5 3.2 7.8 2.1 9.6
Average: 5.64
```

Note: This exercise helps reinforce the understanding of slices by allowing participants to practice creating and manipulating slices, parsing user input, and performing calculations using the values in the slice.