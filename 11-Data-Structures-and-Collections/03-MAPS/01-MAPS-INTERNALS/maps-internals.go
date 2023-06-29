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
	grades := map[string]int{
		"Dan":    85,
		"Emma":   92,
		"Ezra":   78,
		"Ashley": 95,
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
