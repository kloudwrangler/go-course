Exercise: Map Manipulation

Instructions:
1. Create a Go program that demonstrates various map manipulations.
2. Use the knowledge gained from the lesson to complete the following tasks.
3. Run the program and observe the output.
4. Study the provided solution and explanation to understand the concepts.

```go
package main

import "fmt"

func main() {
	// Task 1: Create an empty map called "scores" with string keys and integer values.

	// Task 2: Add the following key-value pairs to the "scores" map:
	// "Alice" -> 85, "Bob" -> 92, "Charlie" -> 78, "Diana" -> 90

	// Task 3: Print the value associated with the key "Charlie" from the "scores" map.

	// Task 4: Update the value of the key "Bob" to 95 in the "scores" map.

	// Task 5: Check if the key "Alice" exists in the "scores" map. Print "Alice exists" if it does, and "Alice doesn't exist" if it doesn't.

	// Task 6: Delete the key-value pair with the key "Charlie" from the "scores" map.

	// Task 7: Iterate over the "scores" map using a for loop and print each key-value pair.

	// Task 8: Print the length of the "scores" map.

	// Task 9: Create a new map called "grades" with integer keys and string values.
	// Initialize it with the following key-value pairs: 85 -> "A", 92 -> "A+", 78 -> "B", 95 -> "A"

	// Task 10: Merge the "grades" map into the "scores" map.

	// Task 11: Print the "scores" map after merging.

}

```

Solution:
```go
package main

import "fmt"

func main() {
	// Task 1: Create an empty map called "scores" with string keys and integer values.
	scores := make(map[string]int)

	// Task 2: Add the following key-value pairs to the "scores" map:
	// "Alice" -> 85, "Bob" -> 92, "Charlie" -> 78, "Diana" -> 90
	scores["Alice"] = 85
	scores["Bob"] = 92
	scores["Charlie"] = 78
	scores["Diana"] = 90

	// Task 3: Print the value associated with the key "Charlie" from the "scores" map.
	fmt.Println("Charlie's score:", scores["Charlie"])

	// Task 4: Update the value of the key "Bob" to 95 in the "scores" map.
	scores["Bob"] = 95

	// Task 5: Check if the key "Alice" exists in the "scores" map.
	// Print "Alice exists" if it does, and "Alice doesn't exist" if it doesn't.
	_, aliceExists := scores["Alice"]
	if aliceExists {
		fmt.Println("Alice exists")
	} else {
		fmt.Println("Alice doesn't exist")
	}

	// Task 6: Delete the key-value pair with the key "Charlie" from the "scores" map.
	delete(scores, "Charlie")

	// Task 7: Iterate over the "scores" map using a for loop and print each key-value pair.
	for key, value := range scores {
		fmt.Println(key, "->", value)
	}

	// Task 8: Print the length of the "scores" map.
	fmt.Println("Number of scores:", len

(scores))

	// Task 9: Create a new map called "grades" with integer keys and string values.
	// Initialize it with the following key-value pairs: 85 -> "A", 92 -> "A+", 78 -> "B", 95 -> "A"
	grades := map[int]string{
		85: "A",
		92: "A+",
		78: "B",
		95: "A",
	}

	// Task 10: Merge the "grades" map into the "scores" map.
	for key, value := range grades {
		scores[key] = value
	}

	// Task 11: Print the "scores" map after merging.
	fmt.Println("Scores after merging:")
	for key, value := range scores {
		fmt.Println(key, "->", value)
	}
}
```

Explanation:
1. The program starts by creating an empty map called "scores" using the `make` function.
2. Key-value pairs are added to the "scores" map using assignment (`=`) and the square brackets (`[]`) syntax.
3. The value associated with the key "Charlie" is printed using `fmt.Println`.
4. The value of the key "Bob" is updated by assigning a new value to it.
5. The existence of the key "Alice" is checked using the two-value assignment (`_, aliceExists := scores["Alice"]`) and the result is printed accordingly.
6. The key-value pair with the key "Charlie" is deleted from the map using the `delete` function.
7. The map is iterated over using a `for` loop, and each key-value pair is printed.
8. The length of the "scores" map is printed using `len(scores)`.
9. A new map called "grades" is created with integer keys and string values, and it is initialized with key-value pairs using a map literal syntax.
10. The "grades" map is merged into the "scores" map by iterating over the "grades" map and assigning each key-value pair to the "scores" map.
11. The final state of the "scores" map after merging is printed by iterating over it using a `for` loop.

This exercise helps reinforce the understanding of map manipulation operations, such as adding, updating, deleting elements, checking key existence, merging maps, and iterating over maps.