Certainly! Here's an example program that utilizes the concepts learned in this lesson about arrays in Go:

```go
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
```

Output:
```
Weekdays:
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Sunday
Third weekday: Wednesday
Original weekdays: [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]
Modified weekdays: [Monday Tuesday Wednesday Thursday Friday Updated Saturday Sunday]
```

In this program, we declare and initialize an array of weekdays. We then use a for loop to iterate over the array and print each weekday. We also demonstrate accessing a specific element by printing the third weekday. Next, we assign the original array to a new array and modify a value in the new array. Finally, we print both the original and modified arrays to see the changes.

Feel free to modify and experiment with the program to further explore the concepts of arrays in Go.