Exercise: Manipulating Maps in Go

Objective: In this exercise, participants will practice manipulating maps in Go by passing them between functions and applying the best practices discussed in the lesson.

Instructions:

1. Create a new Go file named "map_manipulation.go" on your local machine.

2. Define a function named `addToMap` that takes a map of integers to strings as a parameter. The function should add the following key-value pairs to the map:
   - Key: 1, Value: "One"
   - Key: 2, Value: "Two"
   - Key: 3, Value: "Three"

3. Define another function named `modifyMap` that takes the same map of integers to strings as a parameter. The function should modify the value of the key "2" to "New Value".

4. In the `main` function, declare a map of integers to strings.

5. Call the `addToMap` function and pass the map as an argument.

6. Print the contents of the map in the `main` function.

7. Call the `modifyMap` function and pass the map as an argument.

8. Print the contents of the map again.

9. Run the program and observe the output.

Solution:

```go
package main

import "fmt"

func addToMap(m map[int]string) {
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"
}

func modifyMap(m map[int]string) {
	m[2] = "New Value"
}

func main() {
	myMap := make(map[int]string)

	addToMap(myMap)

	fmt.Println("Map after adding values:", myMap)

	modifyMap(myMap)

	fmt.Println("Map after modification:", myMap)
}
```

Explanation:
In this exercise, we create two functions: `addToMap` and `modifyMap`. The `addToMap` function takes a map as a parameter and adds three key-value pairs to the map. The `modifyMap` function also takes a map as a parameter and modifies the value of the key "2" to "New Value". In the `main` function, we declare a map, call the `addToMap` function to populate it, print its contents, call the `modifyMap` function to make changes, and print the contents again.

When you run the program, the output should be:
```
Map after adding values: map[1:One 2:Two 3:Three]
Map after modification: map[1:One 2:New Value 3:Three]
```

This demonstrates the concept of passing maps between functions and modifying them. The first print statement shows the map after adding values, and the second print statement shows the map after the modification. It confirms that the original map was modified within the `modifyMap` function due to the reference-like behavior of maps in Go.