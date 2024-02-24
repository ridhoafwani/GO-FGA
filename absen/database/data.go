package database

type Student struct {
	name     string
	address string
	job string
	reason string
	
}

func GenerateStudents() []Student {
	var students []Student
	students = append(students, Student{name: "John", address: "123 Main St", job: "Teacher", reason: "I like Go"})
	students = append(students, Student{name: "Jane", address: "456 Park Ave", job: "Teacher", reason: "I love Go"})
	students = append(students, Student{name: "Bob", address: "789 Broadway", job: "Teacher", reason: "I dont know"})
	return students
}