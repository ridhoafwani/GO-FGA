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

	fmt.Println(students[i-1])
}

