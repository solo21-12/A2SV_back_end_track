package services

import (
	"example.com/task_three/models"
	"fmt"
)

const (
	BORROWED  = "Borrowed"
	AVAILABLE = "Available"
)

type libraryManagerService struct {
	data          models.Library
	serviceHelper models.ServiceHelper
}

func NewBookService(data models.Library, serviceHelper models.ServiceHelper) models.BookService {
	return &libraryManagerService{
		data:          data,
		serviceHelper: serviceHelper,
	}
}

func (b *libraryManagerService) AddBookService(book models.Book) error {

	_, exists := b.serviceHelper.GetBook(book.ID)

	if exists {
		return fmt.Errorf("the book with the given inforamtion already exists")
	}

	b.data.Books[book.ID] = book
	return nil

}

func (b *libraryManagerService) RemoveBookService(bookID int) error {
	// Check if the book exists
	_, exists := b.serviceHelper.GetBook(bookID)

	if !exists {
		return fmt.Errorf("the book with the given ID does not exist")
	}

	memberID, index := b.serviceHelper.CheckBookBorrowAllMemeber(bookID)

	if memberID != -1 && index != -1 {

		member := b.data.Members[memberID]

		if index >= 0 && index < len(member.BorrowedBooks) {
			member.BorrowedBooks = append(
				member.BorrowedBooks[:index],
				member.BorrowedBooks[index+1:]...,
			)
		} else {
			return fmt.Errorf("index out of range for borrowed books slice")
		}

		b.data.Members[memberID] = member
	}

	delete(b.data.Books, bookID)

	return nil
}

func (b *libraryManagerService) BorrowBookService(bookID int, memberID int) error {
	// check if the book exists
	book, bookExists := b.serviceHelper.GetBook(bookID)

	// check if the member exists
	member, memberExist := b.serviceHelper.GetMember(memberID)

	if !bookExists {
		return fmt.Errorf("the book with the given ID does not exist")
	}

	if !memberExist {
		return fmt.Errorf("the member with the given ID does not exist")
	}

	if borrowed := b.serviceHelper.CheckBorrowed(book); borrowed {
		return fmt.Errorf("the book has already been borrowed")
	}

	book.Status = BORROWED
	b.data.Books[book.ID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	b.data.Members[member.ID] = member

	return nil
}

func (b *libraryManagerService) ReturnBookService(bookID int, memberID int) error {
	// check if the book exists
	book, bookExists := b.serviceHelper.GetBook(bookID)

	// check if the member exists
	member, memberExist := b.serviceHelper.GetMember(memberID)

	if !bookExists {
		return fmt.Errorf("the book with the given ID does not exist")
	}

	if book.Status != BORROWED {
		return fmt.Errorf("the book with the given ID hasn't been borrowed")
	}

	if !memberExist {
		return fmt.Errorf("the member with the given ID does not exist")
	}

	index, borrowed := b.serviceHelper.CheckBookBorrowed(book, member)

	if !borrowed {
		return fmt.Errorf("the given member hasn't borrowed the given book")
	}

	book.Status = AVAILABLE
	b.data.Books[book.ID] = book
	member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)
	b.data.Members[memberID] = member

	return nil
}

func (b *libraryManagerService) ListAvailableBooksService() []models.Book {
	availableBooks := []models.Book{}

	for _, book := range b.data.Books {
		if book.Status == AVAILABLE {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

func (b *libraryManagerService) ListBorrowedBooksService(memberID int) ([]models.Book, error) {
	member, exist := b.serviceHelper.GetMember(memberID)

	if !exist {
		return make([]models.Book, 0), fmt.Errorf("member with the given ID not found")
	}

	return member.BorrowedBooks, nil
}
