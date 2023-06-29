---
marp: true
theme: uncover


---
Title: Multidimensional Slices in Go
Subtitle: Declaring and Composing Slices of Slices

---
## Multidimensional Slices
- Multidimensional slices in Go are slices of slices.
- Each element in the outer slice is itself a slice.

---
Declaring a Multidimensional Slice
- To declare a multidimensional slice, start with the outer slice and initialize each inner slice separately.

---
Code Example - Declaring a Multidimensional Slice:
```go
var matrix [][]int
```

---
Initializing a Multidimensional Slice
- To initialize a multidimensional slice, use the `make` function to create each inner slice.

---
Code Example - Initializing a Multidimensional Slice:
```go
matrix = make([][]int, numRows)
for i := range matrix {
    matrix[i] = make([]int, numCols)
```

---
Accessing and Modifying Elements
- Accessing elements in a multidimensional slice follows the same indexing syntax as regular slices.

---
Code Example - Accessing and Modifying Elements:
```go
matrix[i][j] = newValue
```

---
Composing Slices of Slices
- Composing slices allows you to represent more complex data structures.

---
Code Example - Composing a 2D Grid:
```go
grid := [][]string{
    {"X", "O", "X"},
    {"O", "X", "O"},
    {"X", "O", "X"},
}
```

---
Best Practices
- Use the `len` function to get the length of the outer slice and `len(slice[i])` to get the length of the inner slice.
- Be mindful of memory usage with large and complex multidimensional slices.
- Consider using named types or structs for improved readability and maintainability.

---
Additional Resources
- Official Go Documentation on Slices
- Go by Example: Slices
- "The Go Programming Language" book by Alan A. A. Donovan and Brian W. Kernighan

---
Exercise
- Declare and initialize a tic-tac-toe board using a multidimensional slice.
- Perform operations like printing the board, counting X's and O's, checking for a winner, and flipping the board.

---
Summary
- Multidimensional slices in Go are slices of slices.
- Declare and initialize multidimensional slices using the `make` function.
- Access and modify elements using the same indexing syntax as regular slices.
- Compose slices to create more complex data structures.
- Remember best practices for memory usage and code readability.

---
Questions and Answers

Feel free to customize the slides according to your preferences and teaching style.