package services

import (
	"example.com/task_three/data"
	"example.com/task_three/models"
	"fmt"
)

const (
	BORROWED  = "Borrowed"
	AVAILABLE = "Available"
)

func getBook(bookID int) (models.Book, bool) {
	book, exists := data.OurLibrary.Books[bookID]
	return book, exists
}

func getMember(memberID int) (models.Member, bool) {
	member, exist := data.OurLibrary.Members[memberID]

	return member, exist
}

func checkBookBorrowed(book models.Book, member models.Member) (int, bool) {

	for i, cur_book := range member.BorrowedBooks {
		if cur_book == book {
			return i, true
		}
	}

	return -1, false
}

func AddBook(book models.Book) error {

	_, exists := getBook(book.ID)

	if exists {
		return fmt.Errorf("the book with the given inforamtion already exists")
	}

	data.OurLibrary.Books[book.ID] = book
	return nil

}

func AddMember(member models.Member) error {
	// This method add a new member to our library
	_, exist := getMember(member.ID)

	if exist {
		return fmt.Errorf("the user with the given information already exists")
	}

	data.OurLibrary.Members[member.ID] = member
	return nil

}

func RemoveBook(bookID int) error {
	_, exists := getBook(bookID)

	if exists {
		delete(data.OurLibrary.Books, bookID)
		return nil
	}

	return fmt.Errorf("the book with the given ID does not exist")

}

func BorrowBook(bookID int, memberID int) error {
	// check if the book exists
	book, bookExists := getBook(bookID)

	// check if the member exists
	member, memberExist := getMember(memberID)

	if !bookExists {
		return fmt.Errorf("the book with the given ID does not exist")
	}

	if !memberExist {
		return fmt.Errorf("the member with the given ID does not exist")
	}

	book.Status = BORROWED
	data.OurLibrary.Books[book.ID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	data.OurLibrary.Members[member.ID] = member

	return nil
}

func ReturnBook(bookID int, memberID int) error {
	// check if the book exists
	book, bookExists := getBook(bookID)

	// check if the member exists
	member, memberExist := getMember(memberID)

	if !bookExists {
		return fmt.Errorf("the book with the given ID does not exist")
	}

	if book.Status != BORROWED {
		return fmt.Errorf("the book with the given ID hasn't been borrowed")
	}

	if !memberExist {
		return fmt.Errorf("the member with the given ID does not exist")
	}

	index, borrowed := checkBookBorrowed(book, member)

	if !borrowed {
		return fmt.Errorf("the given member hasn't borrowed the given book")
	}

	book.Status = AVAILABLE
	data.OurLibrary.Books[book.ID] = book
	member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)
	data.OurLibrary.Members[memberID] = member

	return nil
}

func ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}

	for _, book := range data.OurLibrary.Books {
		if book.Status == AVAILABLE {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

func ListBorrowedBooks(memberID int) []models.Book {
	member, exist := getMember(memberID)

	if !exist {
		return make([]models.Book, 0)
	}

	return member.BorrowedBooks
}
