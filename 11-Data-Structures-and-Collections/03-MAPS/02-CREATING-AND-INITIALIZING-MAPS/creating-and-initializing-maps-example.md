# Example: creating and initializing maps in Go:

```go
package main

import (
	"fmt"
)

func main() {
	// Declaring a map using make
	studentScores := make(map[string]int, 100)
	studentScores["Alice"] = 85
	studentScores["Bob"] = 92
	studentScores["Charlie"] = 78

	fmt.Println("Student Scores:", studentScores)

	// Declaring an empty map using a map literal
	employeeDetails := map[string]string{
		"John":   "Software Engineer",
		"Amy":    "Product Manager",
		"Robert": "Data Analyst",
	}

	fmt.Println("Employee Details:", employeeDetails)

	// Declaring a map that stores slices of strings
	attendeesByEvent := map[string][]string{
		"Conference":   []string{"Alice", "Bob", "Charlie"},
		"Workshop":     []string{"David", "Emily", "Frank"},
		"Hackathon":    []string{"George", "Hannah"},
		"Networking":   []string{"Ivy"},
		"Social Event": []string{"Jack", "Kelly"},
	}

	fmt.Println("Attendees by Event:", attendeesByEvent)
}
```

Explanation:
1. We declare a map named `studentScores` using `make` and set an initial capacity of 100. We then add some student names as keys and their corresponding scores as values to the map.
2. We declare a map named `employeeDetails` using a map literal. We provide employee names as keys and their job titles as values.
3. We declare a map named `attendeesByEvent` using a map literal. Each key represents an event name, and the corresponding value is a slice of strings containing the names of attendees for that event.
4. Finally, we print the contents of each map to demonstrate the created and initialized maps.

Output:
```
Student Scores: map[Alice:85 Bob:92 Charlie:78]
Employee Details: map[Amy:Product Manager John:Software Engineer Robert:Data Analyst]
Attendees by Event: map[Conference:[Alice Bob Charlie] Hackathon:[George Hannah] Networking:[Ivy] Social Event:[Jack Kelly] Workshop:[David Emily Frank]]
```

Note: Encourage students to modify the program and experiment with different values to further understand map usage and explore the concepts covered in the lesson.