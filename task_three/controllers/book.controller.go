package controllers

import (
	"bufio"
	"fmt"

	"example.com/task_three/models"
)

const (
	BORROWED  = "Borrowed"
	AVAILABLE = "Available"
)

func separator() {
	fmt.Println("*************************************************")

}

type bookController struct {
	bookService models.BookService
	helper      models.ControllerHelper
}

func NewBookController(bookService models.BookService, helper models.ControllerHelper) models.BookController {
	return &bookController{
		bookService: bookService,
		helper:      helper,
	}
}

func (b *bookController) handleBookTransaction(
	reader *bufio.Reader,
	action func(bookID, memberID int) error,
	actionName string,
) {
	bookID, err := b.helper.GetBookID(reader)
	if err != nil {
		fmt.Println("Error while getting book information:", err)
		return
	}

	memberID, err := b.helper.GetMemberID(reader)
	if err != nil {
		fmt.Println("Error while getting member information:", err)
		return
	}

	if err := action(bookID, memberID); err != nil {
		fmt.Println("Error during", actionName, ":", err)
	} else {
		fmt.Println(actionName, "successfully.")
	}

	separator()
}

func (b *bookController) AddBookController(reader *bufio.Reader) {
	// This method adds a new book to our library

	title, _ := b.helper.GetInput("Book title: ", reader)
	author, _ := b.helper.GetInput("Book author: ", reader)
	cur_ID := b.helper.GenerateNewBookID()

	newBook := models.Book{
		ID:     cur_ID,
		Title:  title,
		Author: author,
		Status: AVAILABLE,
	}

	err := b.bookService.AddBookService(newBook)

	if err != nil {
		fmt.Println("Error adding a book:", err)
	} else {
		fmt.Println("The book has been successfully added")
	}

	separator()

}

func (b *bookController) RemoveBookController(reader *bufio.Reader) {
	// The method removes a borrowed book from a members list adn change the status of the book
	bookID, e := b.helper.GetBookID(reader)

	if e != nil {
		fmt.Printf("Error: %v", e)

	} else {
		err := b.bookService.RemoveBookService(bookID)

		if err != nil {
			fmt.Println("Error removing book:", err)
		} else {
			fmt.Println("Book removed successfully.")
		}

	}

	separator()

}

func (b *bookController) BorrowBookController(reader *bufio.Reader) {
	b.handleBookTransaction(reader, b.bookService.BorrowBookService, "Book borrowed")
}

func (b *bookController) ReturnBookController(reader *bufio.Reader) {
	b.handleBookTransaction(reader, b.bookService.ReturnBookService, "Book returned")
}

func (b *bookController) ListAvailableBooksController() {
	// This method retunr the list of available books
	books := b.bookService.ListAvailableBooksService()

	if len(books) == 0 {
		fmt.Println("No books are available in the library")
		return
	} else {
		fmt.Println("Available Books:")
		separator()
		fmt.Printf("%-10v %-10v %-10v %10v \n", "ID", "Title", "Author", "Status")
		for _, book := range books {
			fmt.Printf("%-10v %-10v %-10v %10v \n", book.ID, book.Title, book.Author, book.Status)
		}

		separator()

	}

	separator()

}

func (b *bookController) ListBorrowedBooksController(reader *bufio.Reader) {
	// This method returns the list of books borrowed by a specifc member
	memberID, err := b.helper.GetMemberID(reader)

	if err != nil {
		fmt.Println("Invalid Member ID:", err)
		return
	}

	books, err := b.bookService.ListBorrowedBooksService(memberID)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else if len(books) == 0 {
		fmt.Println("No books have been borrowed by the given member")
	} else {
		fmt.Printf("Member ID: %-15v", memberID)
		fmt.Println("Borrowed Books:")
		separator()

		fmt.Printf("%-10v %-10v %-10v %10v \n", "ID", "Title", "Author", "Status")
		for _, book := range books {
			fmt.Printf("%-10v %-10v %-10v %10v \n", book.ID, book.Title, book.Author, book.Status)
		}

		separator()

	}

	separator()

}
