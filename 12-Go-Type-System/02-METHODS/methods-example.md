Certainly! Here's a different example that exemplifies methods in Go:

```go
package main

import (
	"fmt"
)

type Student struct {
	name  string
	grade int
}

// Method with a value receiver for Student
func (s Student) PrintInfo() {
	fmt.Printf("Name: %s\nGrade: %d\n", s.name, s.grade)
}

// Method with a pointer receiver for Student
func (s *Student) UpdateGrade(newGrade int) {
	s.grade = newGrade
}

func main() {
	// Create a student
	student := Student{name: "John Doe", grade: 10}

	// Call the PrintInfo method with a value receiver
	student.PrintInfo()

	// Call the UpdateGrade method with a pointer receiver
	student.UpdateGrade(11)

	// Call the PrintInfo method again to see the updated grade
	student.PrintInfo()
}
```

Explanation:
- The program defines a `Student` struct with `name` and `grade` fields.
- The `Student` struct has two methods: `PrintInfo` and `UpdateGrade`.
- The `PrintInfo` method is defined with a value receiver and is used to print the student's name and grade.
- The `UpdateGrade` method is defined with a pointer receiver and is used to update the student's grade.
- In the `main` function, we create an instance of `Student` named `student`.
- We call the `PrintInfo` method on the `student` using a value receiver to print the initial information.
- Then, we call the `UpdateGrade` method on the `student` using a pointer receiver to update the grade.
- Finally, we call the `PrintInfo` method again to see the updated grade.

Expected Output:
```
Name: John Doe
Grade: 10
Name: John Doe
Grade: 11
```

Note: Encourage participants to modify the student's information, add more methods, or experiment with different scenarios to further understand the concept of methods in Go.