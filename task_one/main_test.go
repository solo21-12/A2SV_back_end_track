package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	subjects := []subject{
		{name: "Math", grade: 90},
		{name: "Science", grade: 90},
		{name: "English", grade: 90},
	}

	expectedAverage := 90.0
	actualAverage := calculateAverage(subjects)

	assert.Equal(t, expectedAverage, actualAverage)
}

func TestAcceptScore(t *testing.T) {
	input := "Math\n90\nScience\n85\nEnglish\n88\n"
	reader := bufio.NewReader(strings.NewReader(input))

	subjects := acceptScore(3, reader)
	expectedGrades := []float64{90, 85, 88}

	assert.Equal(t, len(subjects), 3)
	for i, subj := range subjects {
		assert.Equal(t, subj.grade, expectedGrades[i])
	}

}

func TestGetStudent(t *testing.T) {
	input := "Jhon Doe\n3\n"

	reader := bufio.NewReader(strings.NewReader(input))
	name, subject_size := getStudent(reader)

	assert.Equal(t, name, "Jhon Doe")
	assert.Equal(t, subject_size, 3)
}
