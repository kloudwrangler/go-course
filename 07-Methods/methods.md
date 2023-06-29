Lesson: Methods in Go

Topic: Methods

Objective: By the end of this lesson, participants will understand the concept of methods in Go, including how to declare and use methods with value and pointer receivers. They will also learn about choosing between value and pointer receivers and gain insights into how Go determines the relationship between methods and types.

1. Introduction to Methods
   - In Go, a method is a function associated with a type.
   - Methods enable us to define behavior specific to a type.
   - They are declared and defined within the scope of a type.

2. Declaration of a Method with a Value Receiver
   - A method with a value receiver operates on a copy of the value it is called on.
   - It does not modify the original value.
   - Syntax:
     ```go
     func (v Type) methodName() {
         // Method implementation
     }
     ```
   - Example:
     ```go
     type Rectangle struct {
         width, height float64
     }

     func (r Rectangle) Area() float64 {
         return r.width * r.height
     }
     ```

3. Declaration of a Method with a Pointer Receiver
   - A method with a pointer receiver operates directly on the original value.
   - It can modify the value.
   - Syntax:
     ```go
     func (p *Type) methodName() {
         // Method implementation
     }
     ```
   - Example:
     ```go
     type Rectangle struct {
         width, height float64
     }

     func (r *Rectangle) Scale(factor float64) {
         r.width *= factor
         r.height *= factor
     }
     ```

4. Calling a Method from a Variable
   - Methods can be called using the dot notation on a variable of the associated type.
   - The method is invoked based on the receiver type (value or pointer).
   - Example:
     ```go
     rect := Rectangle{width: 10, height: 5}
     area := rect.Area()
     rect.Scale(2)
     ```

5. Choosing Between Pointer vs Value Receiver
   - Use a value receiver when the method does not modify the value.
   - Use a pointer receiver when the method needs to modify the value or avoid unnecessary copying.
   - Consider the semantics and behavior of the type when deciding between value and pointer receivers.
   - Avoid mixing value and pointer receivers for the same method name on a type.

6. Underlying Mechanism of Method Association in Go
   - Go uses automatic conversion between values and pointers to determine the relationship between methods and types.
   - When calling a method with a value receiver on a pointer type, Go automatically dereferences the pointer.
   - When calling a method with a pointer receiver on a value type, Go automatically takes the address of the value.
   - This automatic conversion allows us to call methods conveniently without explicitly dealing with values or pointers.

Additional Resources:
1. Official Go Documentation - Methods: [https://golang.org/doc/effective_go#methods](https://golang.org/doc/effective_go#methods)
2. The Go Programming Language Specification - Method Declarations: [https://golang.org/ref/spec#Method_declarations](https://golang.org/ref/spec#Method_declarations)
3. Tutorialspoint - Go Methods: [https://www.tutorialspoint.com/go/go_methods.htm](https://www.tutorialspoint.com/go/go_methods.htm)
4. Pluralsight - Go: Pointers, Methods, and Interfaces: [https://www.pluralsight.com/courses/go-pointers-methods-interfaces](https://www.pluralsight.com/courses/go-pointers-methods-interfaces)
5. Go by Example - Methods: [

https://gobyexample.com/methods](https://gobyexample.com/methods)

Note: Encourage participants to practice writing and experimenting with methods in Go through coding exercises and examples.