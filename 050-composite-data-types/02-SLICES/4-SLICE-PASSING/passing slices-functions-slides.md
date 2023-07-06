

# Passing Slices between Functions

---

Passing Slices by Value:
- A copy of the slice header is created.
- Modifications to elements inside the function affect the original slice.
- Changes to the length or capacity do not affect the original slice.

```txt
Original Slice: [1, 2, 3, 4, 5]
Modified Slice (Inside Function): [2, 4, 6, 8, 10]
Original Slice (After Function Call): [2, 4, 6, 8, 10]
```

---
Passing Slices by Reference:
- Pass a reference (pointer) to the slice.
- Changes to the length or capacity affect the original slice.
- Assigning a new slice to the parameter does not affect the original slice.

```txt
Original Slice: [1, 2, 3, 4, 5]
Modified Slice (Inside Function): [1, 2, 3, 4, 5, 6]
Original Slice (After Function Call): [1, 2, 3, 4, 5, 6]
```
