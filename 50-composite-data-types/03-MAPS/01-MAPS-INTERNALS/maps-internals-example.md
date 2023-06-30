Certainly! Here's an example program that demonstrates various aspects of maps, including initialization, manipulation, iteration, and best practices:

```go
package main

import "fmt"

func main() {
	// Task 1: Create an empty map called "scores" with string keys and integer values.
	scores := make(map[string]int)

	// Task 2: Add key-value pairs to the "scores" map.
	scores["Alice"] = 85
	scores["Bob"] = 92
	scores["Charlie"] = 78
	scores["Diana"] = 90

	// Task 3: Print the value associated with the key "Charlie" from the "scores" map.
	fmt.Println("Charlie's score:", scores["Charlie"])

	// Task 4: Update the value of the key "Bob" to 95 in the "scores" map.
	scores["Bob"] = 95

	// Task 5: Check if the key "Alice" exists in the "scores" map.
	_, aliceExists := scores["Alice"]
	if aliceExists {
		fmt.Println("Alice exists")
	} else {
		fmt.Println("Alice doesn't exist")
	}

	// Task 6: Delete the key-value pair with the key "Charlie" from the "scores" map.
	delete(scores, "Charlie")

	// Task 7: Iterate over the "scores" map using a for loop and print each key-value pair.
	fmt.Println("Scores:")
	for key, value := range scores {
		fmt.Println(key, "->", value)
	}

	// Task 8: Print the length of the "scores" map.
	fmt.Println("Number of scores:", len(scores))

	// Task 9: Create a new map called "grades" with integer keys and string values.
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

This program showcases the use of maps by performing the following tasks:

1. Creating an empty map called "scores" with string keys and integer values.
2. Adding key-value pairs to the "scores" map.
3. Retrieving and printing the value associated with the key "Charlie".
4. Updating the value of the key "Bob" to 95.
5. Checking the existence of the key "Alice" and printing the result.
6. Deleting the key-value pair with the key "Charlie".
7. Iterating over the "scores" map and printing each key-value pair.
8. Printing the length of the "scores" map.
9. Creating a new map called "grades" with integer keys and string values.
10. Merging the "grades" map into the "scores" map.
11. Printing the "scores" map after merging.

By running this program, you can observe the map manipulation and understand the concepts covered in the lesson.