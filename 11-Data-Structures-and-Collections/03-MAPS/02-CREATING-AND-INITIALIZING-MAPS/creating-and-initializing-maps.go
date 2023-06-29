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
		"Conference":   {"Alice", "Bob", "Charlie"},
		"Workshop":     {"David", "Emily", "Frank"},
		"Hackathon":    {"George", "Hannah"},
		"Networking":   {"Ivy"},
		"Social Event": {"Jack", "Kelly"},
	}

	fmt.Println("Attendees by Event:", attendeesByEvent)
}
