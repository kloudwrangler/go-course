---
marp: true
theme: uncover

---
## Exploring Multidimensional Arrays and Passing Arrays Between Functions in Go

---
Topic: Multidimensional Arrays
```go
// Multidimensional arrays represent tabular or grid-like structures.
// Syntax: var arrayName [rows][columns]Type
// Example:
var matrix [3][3]int
```

---
Accessing Elements of a Two-Dimensional Array
```go
// Elements are accessed using two sets of square brackets.
// Syntax: arrayName[rowIndex][columnIndex]
// Example:
matrix[0][2]
```

---
Assigning Multidimensional Arrays of the Same Type
```go
// Multidimensional arrays can be assigned if they have the same dimensions and element types.
// Example:
source := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
destination := source
```

---
Assigning Multidimensional Arrays by Index
```go
// Specific elements can be assigned in a multidimensional array.
// Example:
array := [2][2]int{{1, 2}, {3, 4}}
array[1][0] = 10
```

---
Passing Arrays Between Functions
```go
// By default, arrays are passed by value.
// Example:
func modifyArray(arr [3]int) {
    // Changes made to arr do not affect the original array
}

array := [3]int{1, 2, 3}
modifyArray(array)
```

---
Passing Arrays Between Functions (contd.)
```go
// To avoid making a copy, arrays can be passed by pointer.
// Example:
func modifyArray(arr *[3]int) {
    // Changes made to arr affect the original array
}

array := [3]int{1, 2, 3}
modifyArray(&array)
```

Note: These code blocks provide a visual representation of the concepts covered in the lesson.