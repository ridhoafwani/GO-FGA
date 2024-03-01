package main

import (
	"absen/database"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]
	students := database.GenerateStudents()

	i, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println("Error converting arg to int:", err)
		return
	}
	if i < 1 || i > len(students) {
		fmt.Println("Invalid student number")
		return
	}
	var student = students[i-1]

	student.Print()
}
