---
marp: true
theme: uncover

---
Working with Maps in Go

---
What are Maps?
- Maps are unordered collections of key-value pairs.
- They provide efficient lookup, insertion, and deletion operations.
- Syntax: `map[keyType]valueType`.

---
Assigning Values to a Map
Code Example:
```go
// Declare and initialize a map
myMap := make(map[string]int)

// Assign values to the map
myMap["apple"] = 3
myMap["banana"] = 5
```
Explanation:
- Declare a map using `make` and provide the key and value types.
- Assign values to the map using the key as an index.

---
Runtime Error with Nil Maps
Code Example:
```go
// Declare a nil map
var myMap map[string]int

// Attempt to assign a value to a nil map
myMap["apple"] = 3 // Runtime error: assignment to entry in nil map
```
Explanation:
- Nil maps have no underlying storage and cannot store values.
- Attempting to assign or retrieve values from a nil map results in a runtime error.
- Always initialize a map before using it using `make` or a map literal.

---
Retrieving Values and Testing Existence
Code Example:
```go
// Retrieve a value from a map
quantity := myMap["apple"]

// Test if a key exists
if _, ok := myMap["banana"]; ok {
    fmt.Println("The key 'banana' exists in the map.")
} else {
    fmt.Println("The key 'banana' does not exist in the map.")
}
```
Explanation:
- Retrieve a value from a map by using the key as an index.
- Use the comma-ok idiom to test if a key exists in the map.
- The value `ok` is `true` if the key exists, otherwise `false`.

---
Retrieving Values and Testing the Value for Existence
Code Example:
```go
// Retrieve a value and check existence
quantity, exists := myMap["apple"]
if exists {
    fmt.Println("The quantity of 'apple' is", quantity)
} else {
    fmt.Println("'apple' does not exist in the map")
}
```
Explanation:
- Use two variables to retrieve a value and check its existence.
- `exists` will be `true` if the key exists, otherwise `false`.

---
Iterating over a Map using for Range
Code Example:
```go
// Iterate over the map using for range
for key, value := range myMap {
    fmt.Println(key, ":", value)
}
```
Explanation:
- Use a for range loop to iterate over a map.
- Each iteration provides a key-value pair which can be accessed using the variables in the loop declaration.

---
Removing an Item from a Map
Code Example:
```go
// Remove an item from the map
delete(myMap, "apple")
```
Explanation:
- Use the `delete` function to remove a key-value pair from the map.
- Provide the map and the key to be deleted as arguments.

---
Best Practices
- Avoid modifying a map while iterating over it.
- Use descriptive key and value types for readability.
- Initialize

 maps using the `make` function.
- Document the expected keys and values in your map.

---
Additional Resources
- Official Go Documentation - Maps
- Effective Go - Maps
- Go by Example - Maps
- Go Maps in Action - A comprehensive guide to working with maps in Go

---
Exercise: Word Frequency Counter
- Create a program that counts the frequency of words in a text file.
- Read the file, split it into words, and store the frequencies in a map.
- Print the word frequency map.

---
Conclusion
- Maps are powerful data structures in Go for key-value storage.
- Practice using maps for efficient data retrieval and manipulation.
- Experiment with maps in your own projects to solidify your understanding.

---
Thank You!
- Questions?

Note: The code examples in the slides provide visual representation of the concepts being discussed, making it easier for participants to understand and follow along. Feel free to customize the code snippets and explanations based on the specific examples and explanations you prefer to use.