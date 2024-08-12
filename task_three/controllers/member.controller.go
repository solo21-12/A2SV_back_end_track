package controllers

import (
	"bufio"
	"fmt"

	"example.com/task_three/models"
)

type memberCotroller struct {
	memberService models.MemberService
	helper                models.ControllerHelper
}

func NewMemberController(memberService models.MemberService, helper models.ControllerHelper) models.MemberController {
	return &memberCotroller{
		memberService: memberService,
		helper:                helper,
	}
}

func (b *memberCotroller) AddMemberController(reader *bufio.Reader) {
	// This method adds a new member to our library

	ID := b.helper.GenerateNewMemberID()
	name, _ := b.helper.GetInput("Member name:", reader)

	newMember := models.Member{
		ID:            ID,
		Name:          name,
		BorrowedBooks: make([]models.Book, 0),
	}

	err := b.memberService.AddMemberService(newMember)
	if err != nil {
		fmt.Println("Error adding member:", err)
	} else {
		fmt.Println("The member has been successfully added")
	}

	separator()

}
