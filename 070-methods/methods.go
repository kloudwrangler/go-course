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
	s.logGradeTransaction()
}

// Private Method -> Users can't use this method directly
func (s Student) logGradeTransaction() {
	fmt.Println("Student updated their grade")
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
