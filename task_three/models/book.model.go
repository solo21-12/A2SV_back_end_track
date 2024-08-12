package models

import "bufio"

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}

type BookService interface {
	AddBookService(book Book) error
	RemoveBookService(bookID int) error
	BorrowBookService(bookID int, memberID int) error
	ReturnBookService(bookID int, memberID int) error
	ListAvailableBooksService() []Book
	ListBorrowedBooksService(memberID int) ([]Book, error)
}

type BookController interface {
	AddBookController(reader *bufio.Reader)
	RemoveBookController(reader *bufio.Reader)
	BorrowBookController(reader *bufio.Reader)
	ReturnBookController(reader *bufio.Reader)
	ListAvailableBooksController()
	ListBorrowedBooksController(reader *bufio.Reader)
}
