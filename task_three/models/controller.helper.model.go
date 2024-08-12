package models

import "bufio"

type ControllerHelper interface {
	GetInput(prompt string, reader *bufio.Reader) (string, error)
	ParseStr(input string) (int, error)
	GenerateNewBookID() int
	GenerateNewMemberID() int
	GetBookID(reader *bufio.Reader) (int, error)
	GetMemberID(reader *bufio.Reader) (int, error)
}
