package models

import "bufio"

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

type MemberService interface {
	AddMemberService(member Member) error
}

type MemberController interface {
	AddMemberController(reader *bufio.Reader)
}
