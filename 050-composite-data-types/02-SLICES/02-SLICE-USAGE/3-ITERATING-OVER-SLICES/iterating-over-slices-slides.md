---
marp: true
theme: uncover

---
# Iterating Over Slices in Go

---
## Iterating Over a Slice using `for range`
- Syntax: `for index, value := range slice`
- `range` provides a copy of each element
- Benefits: simplicity, avoidance of off-by-one errors

---
Example 1: Iterating Over a Slice using `for range`
```
intSlice := []int{10, 20, 30, 40, 50}
for _, value := range intSlice {
    fmt.Println(value)
}
```
Output:
```
10
20
30
40
50
```

---
## Using the Blank Identifier to Ignore Index Value
- Syntax: `for _, value := range slice`
- `_` signals intentional unused index value

---
Example 2: Using the Blank Identifier to Ignore Index Value
```
strSlice := []string{"apple", "banana", "cherry", "date"}
for index, value := range strSlice {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}
```
Output:
```
Index: 0, Value: apple
Index: 1, Value: banana
Index: 2, Value: cherry
Index: 3, Value: date
```

---
## Iterating Over a Slice using a Traditional `for` Loop
- Syntax: `for i := 0; i < len(slice); i++`
- Provides fine-grained control over the iteration

---
Example 3: Iterating Over a Slice using a Traditional `for` Loop
```
floatSlice := []float64{3.14, 1.618, 2.718, 0.577}
sum := 0.0
for i := 0; i < len(floatSlice); i++ {
    sum += floatSlice[i]
}
fmt.Println("Sum:", sum)
```
Output:
```
Sum: 7.027
```

---
Additional Resources:
- Official Go documentation on slices
- A Tour of Go - Loops and Functions
- The Go Programming Language Specification
- "Effective Go" guide

---
Summary:
- Iterating over slices is a common operation in Go programming.
- `for range` provides a simple and safe way to iterate over slices, avoiding off-by-one errors.
- The blank identifier `_` can be used to ignore the index value when necessary.
- Traditional `for` loops offer more control, especially when accessing elements by index.
- Practice and experimentation will further solidify your understanding.

Note: The number of slides may vary based on the presentation style and pace.