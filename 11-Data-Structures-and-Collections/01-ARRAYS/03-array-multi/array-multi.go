package main

import "fmt"

func main() {
	// Declare and initialize an array of weekdays
	var weekdays = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Print the weekdays using a for loop
	fmt.Println("Weekdays:")
	for i := 0; i < len(weekdays); i++ {
		fmt.Println(weekdays[i])
	}

	// Access and print a specific weekday
	fmt.Println("Third weekday:", weekdays[2])

	// Assign one array to another
	var newWeekdays = weekdays

	// Modify a value in the newWeekdays array
	newWeekdays[5] = "Updated Saturday"

	// Print the original and modified arrays
	fmt.Println("Original weekdays:", weekdays)
	fmt.Println("Modified weekdays:", newWeekdays)
}
