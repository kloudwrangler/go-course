---
marp: true
theme: uncover
---
## Working with Slices

---
Title: Assigning and Slicing

---
Code Block:
```
// Declaring an array using an array literal
arr := [5]int{1, 2, 3, 4, 5}
```
Explanation: 
- Arrays have a fixed length in Go.
- Array literals allow us to declare an array with initial values.

---
Code Block:
```
// Taking the slice of a slice
slice := arr[2:4]
```
Explanation:
- Slicing creates a new slice from an existing slice or array.
- The resulting slice includes elements from the specified range.

---
Code Block:
```
// How length and capacity are calculated
length := len(slice)
capacity := cap(slice)
```
Explanation:
- `len(slice)` returns the number of elements in the slice.
- `cap(slice)` returns the maximum number of elements the slice can hold.

---
Code Block:
```
// Calculating the new length and capacity
newSlice := slice[1:3]
newLength := len(newSlice)
newCapacity := cap(newSlice)
```
Explanation:
- Taking a slice of a slice can change the length and capacity.
- The new length is determined by the range of elements in the new slice.
- The new capacity is calculated based on the original slice's capacity.

---
Code Block:
```
// Potential consequence of making changes to a slice
slice[0] = 10
```
Explanation:
- Slices in Go are reference types.
- Modifying a slice can affect other slices sharing the same underlying array.

---
Code Block:
```
// Runtime error showing index out of range
element := slice[10]
```
Explanation:
- Accessing an index that is out of range will result in a runtime error.
- Go will panic and display an "index out of range" error message.

---
Title: Additional Resources

---
- Official Go Documentation - Slices
- A Tour of Go - Slices
- Effective Go - Slices
- Go by Example - Slices
- SliceTricks

Note: Feel free to customize the design and layout of the slides to suit your presentation style.