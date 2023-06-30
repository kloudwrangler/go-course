---
marp: true
theme: uncover

---
Title: Creating and Initializing Maps in Go

---
## Introduction to Maps
- Maps in Go are key-value pairs data structures.
- Efficient for lookups and updates based on a unique key.
- In this lesson, we'll learn how to create and initialize maps in Go.

---
## Declaring a Map Using Make
- Syntax: `mapName := make(map[keyType]valueType, initialCapacity)`
- Example:
```
studentScores := make(map[string]int, 100)
```
- Creates a map named `studentScores` with a capacity of 100.

---
## Declaring an Empty Map Using a Map Literal
- Syntax: `mapName := map[keyType]valueType{}`
- Example:
```
employeeDetails := map[string]string{}
```
- Creates an empty map named `employeeDetails`.

---
## Declaring a Map That Stores Slices of Strings
- Syntax: `mapName := map[keyType][]valueType{}`
- Example:
```
attendeesByEvent := map[string][]string{}
```
- Creates a map named `attendeesByEvent` that stores slices of strings.

---
## Best Practices
- Always initialize a map before using it to avoid nil pointer dereference errors.
- Use the `make` function to allocate memory for the map if you know the expected size or capacity in advance.
- Use map literals for small maps or when the initial capacity is not known.
- Ensure the key type is comparable to be used as a map key.

---
## Additional Resources
- Official Go Documentation on Maps: https://golang.org/doc/effective_go#maps
- A Tour of Go: Maps - https://tour.golang.org/moretypes/19
- "Working with Maps in Go" - Tutorial by Callicoder: https://www.callicoder.com/golang-maps/
- "Go Maps in Action" - Blog Post by Dave Cheney: https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics

---
Summary:
- Declaring a map using `make` allows setting an initial capacity.
- An empty map can be declared using a map literal.
- Maps can store values of any type, including slices.
- Follow best practices to ensure proper map usage.

---
Conclusion:
- In this lesson, we covered creating and initializing maps in Go.
- You learned about using `make`, map literals, and storing slices in maps.
- Remember to practice the best practices for working with maps.
- Explore the additional resources provided for further learning.

---
Thank You!
- Questions?