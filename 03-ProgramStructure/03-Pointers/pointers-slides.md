---
marp: true
theme: uncover

---

---
## Pointers in Go
- Welcome to the lesson on pointers in Go!

---
## Understanding Pointers
- Pointers allow us to directly manipulate memory.
- Pointers hold the memory address of another variable.

---
## Declaring and Initializing Pointers
- Syntax: `var variableName *dataType`
  - Example: `var ptr *int`
- Initialization: Use the `&` operator to assign the address of a variable to a pointer.
  - Example: `ptr = &num`

---
## Dereferencing Pointers
- Dereferencing: Accessing the value stored at the memory address a pointer points to.
- Use the `*` operator to dereference a pointer.
  - Example: `*ptr`

---
## Passing Pointers to Functions
- Functions can accept pointers as parameters.
- Allows modification of the original variable.
  - Example function:
    ```go
    func increment(ptr *int) {
      *ptr = *ptr + 1
    }
    ```

---
## Best Practices for Using Pointers
- Use pointers when necessary, such as modifying variable values or working with large data structures.
- Avoid unnecessary complexity by minimizing pointer usage.
- Be cautious with pointer arithmetic.
- Document ownership and pointer usage.

---
## Exercise
- Declare and initialize pointers.
- Dereference and modify pointers.
- Pass pointers to functions.

---
## Additional Resources
- Go Documentation on Pointers: [https://golang.org/doc/faq#pointers](https://golang.org/doc/faq#pointers)
- A Tour of Go - Pointers: [https://tour.golang.org/moretypes/1](https://tour.golang.org/moretypes/1)
- Mastering Go - Pointers: Book by Mihalis Tsoukalos
- Learn Go with Tests - Pointers: [https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors)

---
## Conclusion
- Pointers are powerful for memory manipulation in Go.
- Understanding declaration, initialization, dereferencing, and passing pointers to functions is essential.
- Use pointers judiciously and follow best practices for code readability and maintainability.

Note: The code examples provided in the slides should be displayed in code blocks or on-screen as readable code snippets for students to understand and follow along with during the lesson.