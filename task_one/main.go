package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type subject struct {
	name  string
	grade float64
}

type student struct {
	subject_size        int
	subjects            []subject
	total_avarage_grade float64
}

var students = make(map[string]student)

func getInput(promt string, reader *bufio.Reader) (string, error) {
	fmt.Print(promt)

	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func getStudent(reader *bufio.Reader) (string, int) {

	name, err := getInput("What's your name: ", reader)

	if err != nil {
		fmt.Println("The entered name isn't valid")
		return getStudent(reader)
	}

	var subject_size int

	for {
		input, _ := getInput("How many subjects have you taken: ", reader)
		s, err := strconv.Atoi(input)
		if err != nil || s <= 0 {
			fmt.Println("Invalid number. Please enter a positive integer.")
		} else {
			subject_size = s
			break
		}
	}

	return name, subject_size

}

func acceptScore(subject_size int, reader *bufio.Reader) []subject {

	cur_subjects := []subject{}

	for i := 0; i < subject_size; i++ {
		name, _ := getInput("Subject name: ", reader)

		var grade float64

		for {
			input, _ := getInput("Subject grade: ", reader)
			g, _ := strconv.ParseFloat(input, 64)

			if g < 0 || g > 100 {
				fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
			} else {
				grade = g
				break
			}

		}

		var cur_subj subject = subject{name: name, grade: grade}

		cur_subjects = append(cur_subjects, cur_subj)
	}

	return cur_subjects

}

func calculateAverage(cur_subject []subject) float64 {
	average := 0.0

	for _, cur := range cur_subject {
		average += cur.grade
	}

	return average / float64(len(cur_subject))
}

func printStudent(name string) {
	cur_student := students[name]

	fmt.Println("*********************************************")
	fmt.Println("The provided students information is as follows")
	fmt.Println("*********************************************")
	fmt.Printf("Student name: %v \n", name)

	fmt.Printf("Total number of subjects the student has taken: %v \n", cur_student.subject_size)
	fmt.Println("*********************************************")

	fmt.Printf("%-25v Score \n", "Subject name")

	for i := 0; i < len(cur_student.subjects); i++ {
		fmt.Printf("%-25v %v \n", cur_student.subjects[i].name, cur_student.subjects[i].grade)
	}

	fmt.Printf("%-25v %v \n", "Average score", cur_student.total_avarage_grade)

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name, subject_size := getStudent(reader)
	subjects := acceptScore(subject_size, reader)
	average := calculateAverage(subjects)

	students[name] = student{subject_size: subject_size, subjects: subjects, total_avarage_grade: average}

	printStudent(name)

}
