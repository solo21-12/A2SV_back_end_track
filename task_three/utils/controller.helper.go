package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"example.com/task_three/models"
)

type controllerHelper struct {
	data models.Library
}

func NewControllerHelper(data models.Library) models.ControllerHelper {
	return &controllerHelper{
		data: data,
	}
}

func (h *controllerHelper) GetInput(prompt string, reader *bufio.Reader) (string, error) {
	// This methods accepts users input
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), err
}

func (h *controllerHelper) ParseStr(input string) (int, error) {
	// This method parses the given string to int
	res, err := strconv.Atoi(input)

	return res, err

}

func (h *controllerHelper) GenerateNewBookID() int {
	// This method generates a new book id
	return len(h.data.Books) + 1
}

func (h *controllerHelper) GenerateNewMemberID() int {
	// This method generates a new member id
	return len(h.data.Members) + 1
}

func (h *controllerHelper) GetBookID(reader *bufio.Reader) (int, error) {
	// This method accepts book's id
	bid, _ := h.GetInput("Enter Book ID: ", reader)
	bookID, bErr := h.ParseStr(bid)

	if bErr != nil {
		fmt.Println("Invalid Book ID:", bErr)
		return -1, bErr
	}

	return bookID, nil

}

func (h *controllerHelper) GetMemberID(reader *bufio.Reader) (int, error) {
	// The method accepts member's id
	mid, _ := h.GetInput("Enter Member ID: ", reader)

	memberID, mErr := h.ParseStr(mid)

	if mErr != nil {
		fmt.Println("Invalid Member ID:", mErr)
		return -1, mErr
	}

	return memberID, nil

}


