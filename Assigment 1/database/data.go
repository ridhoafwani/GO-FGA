package database

import "fmt"

type Student struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func (s Student) Print() {
	fmt.Println("Nama:", s.Name)
	fmt.Println("Address:", s.Address)
	fmt.Println("Job:", s.Job)
	fmt.Println("Reason:", s.Reason)
}

func GenerateStudents() []Student {
	var students []Student
	students = append(students, Student{Name: "John", Address: "123 Main St", Job: "Teacher", Reason: "I like Go"})
	students = append(students, Student{Name: "Jane", Address: "456 Park Ave", Job: "Teacher", Reason: "I love Go"})
	students = append(students, Student{Name: "Bob", Address: "789 Broadway", Job: "Teacher", Reason: "I dont know"})
	return students
}
