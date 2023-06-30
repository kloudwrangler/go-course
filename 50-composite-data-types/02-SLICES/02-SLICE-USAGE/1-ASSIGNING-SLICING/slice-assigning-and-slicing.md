Lesson: Working with Slices

1. Assigning and Slicing

- Declaring an array using an array literal:
  - In Go, arrays have a fixed length that needs to be specified at compile-time. To declare an array using an array literal, you can use the following syntax:
    ```go
    arr := [5]int{1, 2, 3, 4, 5}
    ```
    In this example, we declare an array named `arr` with a length of 5 and initialize it with the values 1, 2, 3, 4, and 5.

- Taking the slice of a slice:
  - Slicing allows you to create a new slice by specifying a range of elements from an existing slice or array. To take a slice of a slice, you can use the following syntax:
    ```go
    slice := arr[start:end]
    ```
    The resulting `slice` will contain elements from the `start` index up to, but not including, the `end` index.

- How length and capacity are calculated:
  - A slice in Go is a dynamically-sized, flexible view into the elements of an array. It consists of a pointer to the underlying array, the length of the slice, and the capacity. The length represents the number of elements in the slice, and the capacity is the maximum number of elements the slice can hold. When you create a slice, the length and capacity are calculated based on the provided range.

- Calculating the new length and capacity:
  - When you take a slice of a slice or modify an existing slice, the length and capacity of the resulting slice may change. To calculate the new length and capacity of a slice, you can use the following formulas:
    - New length: `end - start`
    - New capacity: `original capacity - start`

- Potential consequence of making changes to a slice:
  - Slices in Go are reference types, which means they point to an underlying array. If you make changes to a slice, such as modifying its elements or resizing it, it may affect other slices that share the same underlying array. This behavior can lead to unexpected consequences if you're not aware of it.

- Runtime error showing index out of range:
  - One common error when working with slices is accessing an element using an index that is out of range. This can happen if you try to access an index that is less than 0 or greater than or equal to the length of the slice. When such an error occurs, Go will panic at runtime and display an "index out of range" error message.

Additional Resources:
1. Official Go Documentation - Slices: [https://golang.org/doc/slices](https://golang.org/doc/slices)
2. A Tour of Go - Slices: [https://tour.golang.org/moretypes/7](https://tour.golang.org/moretypes/7)
3. Effective Go - Slices: [https://golang.org/doc/effective_go#slices](https://golang.org/doc/effective_go#slices)
4. Go by Example - Slices: [https://gobyexample.com/slices](https://gobyexample.com/slices)
5. SliceTricks: [https://github.com/golang/go/wiki/SliceTricks](https://github.com/golang/go/wiki/SliceTricks)

Note: It's recommended to provide hands-on coding exercises to the participants to reinforce the concepts discussed in the lesson.