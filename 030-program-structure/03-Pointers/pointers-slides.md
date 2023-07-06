
## Pointers in Go
- Welcome to the lesson on pointers in Go!

---
## Understanding Pointers
- Pointers allow us to directly manipulate memory.
- Pointers hold the memory address of another variable.

---
## Declaring Pointers

- Syntax
	- `var variableName *dataType`
- Example

```go[6]
func increment(ptr *int) {
	*ptr = *ptr + 1
}

func main() {
	var ptr *int // Declare
	num := 42
	ptr = &num // Initialize
	fmt.Println("Value of num:", *ptr) // Derefrencing
	*ptr = 99 // Derefrencing for changing value
	increment(ptr) // Pass to a function
	fmt.Println("Value of num after increment:", num)
}
```
---
## Intializing Pointers

- Use the `&` operator to assign the address of a variable to a pointer.

```go[6,8]
func increment(ptr *int) {
	*ptr = *ptr + 1
}

func main() {
	var ptr *int // Declare
	num := 42
	ptr = &num // Initialize
	fmt.Println("Value of num:", *ptr) // Derefrencing
	*ptr = 99 // Derefrencing for changing value
	increment(ptr) // Pass to a function
	fmt.Println("Value of num after increment:", num)
}
```

---
## Dereferencing Pointers
- Dereferencing: Accessing the value stored at the memory address a pointer points to.
- Use the `*` operator to dereference a pointer.
  - Example: `*ptr`

```go[9]
func increment(ptr *int) {
	*ptr = *ptr + 1
}

func main() {
	var ptr *int // Declare
	num := 42
	ptr = &num // Initialize
	fmt.Println("Value of num:", *ptr) // Derefrencing
	*ptr = 99 // Derefrencing for changing value
	increment(ptr) // Pass to a function
	fmt.Println("Value of num after increment:", num)
}
```

---
## Passing Pointers to Functions
- Functions can accept pointers as parameters.
- Allows modification of the original variable.

```go[1-3,11-12]
func increment(ptr *int) {
	*ptr = *ptr + 1
}

func main() {
	var ptr *int // Declare
	num := 42
	ptr = &num // Initialize
	fmt.Println("Value of num:", *ptr) // Derefrencing
	*ptr = 99 // Derefrencing for changing value
	increment(ptr) // Pass to a function
	fmt.Println("Value of num after increment:", num)
}
```

---
## Best Practices for Using Pointers
- Use pointers when necessary, such as modifying variable values or working with large data structures.
- Avoid unnecessary complexity by minimizing pointer usage.
- Be cautious with pointer arithmetic.
