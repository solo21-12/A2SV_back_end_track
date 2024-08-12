package services

import (
	"fmt"

	"example.com/task_three/models"
)

type memberService struct {
	data          models.Library
	serviceHelper models.ServiceHelper
}

func NewMemberService(data models.Library, serviceHelper models.ServiceHelper) models.MemberService {
	return &memberService{
		data:          data,
		serviceHelper: serviceHelper,
	}
}

func (b *memberService) AddMemberService(member models.Member) error {
	// This method add a new member to our library
	_, exist := b.serviceHelper.GetMember(member.ID)

	if exist {
		return fmt.Errorf("the user with the given information already exists")
	}

	b.data.Members[member.ID] = member
	return nil

}
