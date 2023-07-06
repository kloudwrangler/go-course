---
marp: true
theme: uncover

Apologies for the oversight. Here are the revised slides with code examples included:

---
Title: Passing Maps Between Functions

---
Objective:
- Understand how to pass maps between functions in Go.
- Apply best practices for map manipulation.

---
Introduction to Maps:
- Maps are unordered collections of key-value pairs.
- Maps are reference types and can be modified directly.

---
Passing Maps by Value vs. Reference:
- Passing by value creates a copy of the map.
- Passing by reference allows modifications to the original map.

---
Passing Maps by Value:
- Isolates modifications within the function.
- Prevents unintended changes to the original map.

Code Example:
```go
func addToMap(m map[int]string) {
    m[1] = "One"
    m[2] = "Two"
}

myMap := make(map[int]string)
addToMap(myMap)
fmt.Println(myMap) // Output: map[1:One 2:Two]
```

---
Passing Maps by Reference:
- Allows functions to modify the original map.
- Consider concurrency and synchronization for shared maps.

Code Example:
```go
func modifyMap(m map[int]string) {
    m[2] = "New Value"
}

myMap := make(map[int]string)
modifyMap(myMap)
fmt.Println(myMap) // Output: map[2:New Value]
```

---
Best Practices:
- Check if a key exists before accessing its value.
- Use the comma-ok idiom for handling missing keys.
- Consider concurrency and synchronization for shared maps.
- Document the expected structure of maps.
- Preallocate the map with an appropriate size to avoid excessive growth.

---
Additional Resources:
- Official Go documentation on maps: https://golang.org/doc/effective_go#maps
- "Go Maps in Action" by William Kennedy: https://www.ardanlabs.com/blog/2013/12/maps-in-go-look-inside.html
- "The Go Programming Language" book by Alan A. A. Donovan and Brian W. Kernighan: Chapter 7 - Maps
- "Go by Example: Maps" tutorial: https://gobyexample.com/maps

---
Conclusion:
- Recap the key points covered in the lesson.
- Encourage participants to practice passing maps between functions and apply the best practices discussed.
- Answer any questions and provide clarifications as needed.

Note: The code examples provided are for reference and can be used during the presentation to demonstrate the concepts.