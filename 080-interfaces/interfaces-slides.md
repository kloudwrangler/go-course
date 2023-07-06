---
marp: true
theme: uncover

---

# Interfaces in Go Programming
### Leveraging Polymorphism and Code Reusability

---
## Introduction to Interfaces
- Interfaces define behavior without specifying the underlying type.
- They enable code reuse and flexibility in Go programs.
- Different types can be treated interchangeably if they implement the same set of methods.

---
## Example Program - Multiple Interfaces

Code Block:
```go
type Writer interface {
    Write(data []byte) (int, error)
}

type Closer interface {
    Close() error
}

type FileWriter struct {
    // Fields and implementations
}

func (fw FileWriter) Write(data []byte) (int, error) {
    // Write implementation
}

func (fw FileWriter) Close() error {
    // Close implementation
}

func main() {
    var w Writer
    var c Closer

    file := FileWriter{}
    w = file
    c = file

    // File can be treated as both Writer and Closer
}
```

Explanation:
- This code demonstrates how a single type, `FileWriter`, can implement multiple interfaces: `Writer` and `Closer`.
- By implementing the required methods for each interface, `FileWriter` can be used interchangeably as a `Writer` or `Closer`.

---
## Implementation of Interfaces

Code Block:
```
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}

func main() {
    rect := Rectangle{width: 5, height: 3}
    var shape Shape
    shape = rect

    // Rectangle implements the Shape interface
}
```

Explanation:
- This code demonstrates the implementation of the `Shape` interface for the `Rectangle` struct.
- The `Rectangle` struct defines the `Area()` and `Perimeter()` methods required by the `Shape` interface.
- By assigning a `Rectangle` instance to a variable of type `Shape`, we can treat the `Rectangle` as a `Shape`.

---
## Method Sets

Code Block:
```
type Shape interface {
    Area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r *Rectangle) Resize(factor float64) {
    r.width *= factor
    r.height *= factor
}

func main() {
    rect := Rectangle{width: 5, height: 3}
    var shape Shape

    shape = rect // Valid, method set of Rectangle includes value receivers

    rectPtr := &rect
    shape = rectPtr // Valid, method set of *Rectangle includes pointer receivers
}
```

Explanation:
- This code showcases method sets in Go and their impact on interface implementation and type conversions.
- The `Rectangle` struct defines the `Area()` method with a value receiver, and the `Resize()` method with a pointer receiver.
- Assigning a `Rectangle` instance to a variable of type `Shape` is valid because the `Area()` method has a value receiver.
- Assigning a `*Rectangle` (pointer to `Rectangle`) to a variable of type `Shape` is also valid because the `Area()` method has a pointer receiver.

---
## Polymorphism through Type Embedding



Code Block:
```
type Animal interface {
    Speak() string
}

type Dog struct {
    // Fields and implementations
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
    // Fields and implementations
}

func (c Cat) Speak() string {
    return "Meow!"
}

type AnimalSpeaker struct {
    Animal
}

func main() {
    dog := Dog{}
    cat := Cat{}

    speaker := AnimalSpeaker{dog}
    fmt.Println(speaker.Speak()) // Output: Woof!

    speaker.Animal = cat
    fmt.Println(speaker.Speak()) // Output: Meow!
}
```

Explanation:
- This code demonstrates how type embedding can enable polymorphism.
- The `Dog` and `Cat` types implement the `Animal` interface with their respective `Speak()` methods.
- The `AnimalSpeaker` struct embeds the `Animal` interface, allowing it to hold any value that implements `Animal`.
- By assigning a `Dog` instance to the `AnimalSpeaker`, we can call the `Speak()` method and get "Woof!" as the output.
- Similarly, by assigning a `Cat` instance to the `AnimalSpeaker`, we can call the `Speak()` method and get "Meow!" as the output.

---
Summary:
- Interfaces provide a way to define behavior without specifying the underlying type.
- Implementing interfaces allows types to be treated interchangeably based on shared methods.
- Method sets determine the behavior of methods based on value or pointer receivers.
- Polymorphism can be achieved through type embedding, enabling multiple behaviors for a single type.

---
Additional Resources:
- [Go Interface Tour](https://go.dev/tour/methods/15)
- [Effective Go - Interfaces and other Types](https://golang.org/doc/effective_go.html#interfaces_and_types)



---
Title: Interfaces in Go Programming

---
Objective:
- Understand the concept of interfaces in Go programming language.
- Learn how a type can implement multiple interfaces.
- Explore method sets and polymorphism in Go.
- Gain knowledge about type embedding and its applications.
- Learn about struct and interface embedding.
- Understand type assertions and type switches for working with interfaces.

---
What are Interfaces?
- An interface is a collection of method signatures.
- In Go, interfaces define behavior, not structure.
- A type implements an interface by implementing its methods.

---
Example Program: Multiple Interfaces
- Demonstrate how a type can implement multiple interfaces.
- Show code examples where a single type is considered different types based on the interfaces it implements.
  
```go
type Shape interface {
    Area() float64
}

type Perimeter interface {
    Perimeter() float64
}

type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    var s Shape
    s = rect
    fmt.Println(s.Area())      // Output: 50
    var p Perimeter
    p = rect
    fmt.Println(p.Perimeter()) // Output: 30
}
```

---
Implementation of Interfaces
- Discuss how to implement an interface in Go.
- Explain that a type implicitly implements an interface if it contains all the methods defined in the interface.

---
Method Sets
- Explain the concept of method sets.
- Method set of a type determines the interface it implements.
- Discuss value receiver vs. pointer receiver and their impact on method sets.
  
```go
type MyInt int

func (m MyInt) Print() {
    fmt.Println(m)
}

func (m *MyInt) Increment() {
    *m++
}

func main() {
    var i MyInt = 5
    i.Print()        // Output: 5
    var p *MyInt = &i
    p.Increment()
    p.Print()        // Output: 6
}
```

---
Polymorphism in Go
- Explain how interfaces enable polymorphism in Go.
- Code example showcasing polymorphic behavior using interfaces.

```go
type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{Dog{}, Cat{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

---
Type Embedding
- Introduce type embedding as a way to reuse code and extend behavior.
- Explain that a struct can embed fields from other structs.
- Show code examples of struct embedding in Go.

```go
type Person struct {
    name string
    age  int
}

type Employee struct {
    Person
    role string
}

func main() {
    emp := Employee{
        Person: Person{
            name: "John",
            age:  30,
        },
        role: "Developer",
    }
    fmt.Println(emp.name) // Output: John
    fmt.Println(emp.age)  // Output: 30
    fmt.Println(emp.role) // Output: Developer
}
```



---
Interface Embedding
- Discuss interface embedding as a way to compose interfaces.
- Explain that an interface can embed other interfaces.
- Show code examples of interface embedding in Go.

```go
type Shape interface {
    Area() float64
}

type Perimeter interface {
    Perimeter() float64
}

type Geometry interface {
    Shape
    Perimeter
}

type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    var g Geometry = rect
    fmt.Println(g.Area())       // Output: 50
    fmt.Println(g.Perimeter())  // Output: 30
}
```

---
Type Assertions
- Introduce type assertions for working with interfaces.
- Explain how type assertions allow us to extract the underlying concrete type from an interface.
- Provide code examples of type assertions in Go.

---
Type Switches
- Discuss type switches as an alternative to type assertions.
- Explain how type switches enable branching based on the type of an interface value.
- Show code examples of type switches in Go.

---
Additional Resources
- Provide links to additional resources for further learning:
  - Go Tour: Methods - Type Assertions (https://go.dev/tour/methods/15)
  - Go Tour: Methods - Type Switches (https://go.dev/tour/methods/16)

---
Summary
- Recap the key points covered in the lesson.
- Emphasize the importance of interfaces in Go programming.
- Encourage participants to explore and experiment with interfaces in their own code.

---
Thank You!
- Conclude the lesson and thank the participants for their attention.
- Provide contact information for any further questions or assistance.

Note: These slides include code snippets to aid in visualization. Feel free to modify or add more code examples as per your requirements.