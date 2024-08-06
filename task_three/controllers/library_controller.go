package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"example.com/task_three/data"
	"example.com/task_three/models"
	"example.com/task_three/services"
)

const (
	BORROWED  = "Borrowed"
	AVAILABLE = "Available"
)

func separetor() {
	fmt.Println("*************************************************")

}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	// This methods accepts users input
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), err
}

func parseStr(input string) (int, error) {
	// This method parses the given string to int
	res, err := strconv.Atoi(input)

	return res, err

}

func generateNewBookID() int {
	// This method generates a new book id
	return len(data.OurLibrary.Books) + 1
}

func generateNewMemberID() int {
	// This method generates a new member id
	return len(data.OurLibrary.Members) + 1
}

func getBookID(reader *bufio.Reader) (int, error) {
	// This method accepts book's id
	bid, _ := getInput("Enter Book ID: ", reader)
	bookID, bErr := parseStr(bid)

	if bErr != nil {
		fmt.Println("Invalid Book ID:", bErr)
		return -1, bErr
	}

	return bookID, nil

}

func getMemberID(reader *bufio.Reader) (int, error) {
	// The method accepts member's id
	mid, _ := getInput("Enter Member ID: ", reader)

	memberID, mErr := parseStr(mid)

	if mErr != nil {
		fmt.Println("Invalid Member ID:", mErr)
		return -1, mErr
	}

	return memberID, nil

}

func AddBookController(reader *bufio.Reader) {
	// This method adds a new book to our library

	title, _ := getInput("Book title: ", reader)
	author, _ := getInput("Book author: ", reader)
	cur_ID := generateNewBookID()

	newBook := models.Book{
		ID:     cur_ID,
		Title:  title,
		Author: author,
		Status: AVAILABLE,
	}

	err := services.AddBook(newBook)

	if err != nil {
		fmt.Println("Error adding a book:", err)
	} else {
		fmt.Println("The book has been successfully added")
	}

	separetor()

}

func AddMemberController(reader *bufio.Reader) {
	// This method adds a new member to our library

	ID := generateNewMemberID()
	name, _ := getInput("Member name:", reader)

	newMember := models.Member{
		ID:            ID,
		Name:          name,
		BorrowedBooks: make([]models.Book, 0),
	}

	err := services.AddMember(newMember)
	if err != nil {
		fmt.Println("Error adding member:", err)
	} else {
		fmt.Println("The member has been successfully added")
	}

	separetor()

}

func RemoveBookController(reader *bufio.Reader) {
	// The method removes a borrowed book from a members list adn change the status of the book
	bookID, e := getBookID(reader)

	if e != nil {
		fmt.Printf("Error: %v", e)

	} else {
		err := services.RemoveBook(bookID)

		if err != nil {
			fmt.Println("Error removing book:", err)
		} else {
			fmt.Println("Book removed successfully.")
		}

	}

	separetor()

}

func BorrowBookController(reader *bufio.Reader) {
	// This method handles the borrowing of a book
	bookID, e := getBookID(reader)
	memberID, em := getMemberID(reader)

	if e != nil {
		fmt.Println("Error while getting information", e)

	} else if em != nil {
		fmt.Println("Error while getting information", em)
	} else {
		err := services.BorrowBook(bookID, memberID)

		if err != nil {
			fmt.Println("Error borrowing book: ", err)
		} else {
			fmt.Println("Book borrowed successfully.")
		}
	}

	separetor()

}

func ReturnBookController(reader *bufio.Reader) {
	// This method handles the returning of a book

	bookID, e := getBookID(reader)
	memberID, em := getMemberID(reader)

	if e != nil {
		fmt.Println("Error while getting information", e)

	} else if em != nil {
		fmt.Println("Error while getting information", em)
	} else {
		err := services.ReturnBook(bookID, memberID)

		if err != nil {
			fmt.Println("Error returning book: ", err)
		} else {
			fmt.Println("Book returned successfully.")
		}
	}

	separetor()

}

func ListAvailableBooksController() {
	// This method retunr the list of available books
	books := services.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("No books have been added to the library")
		return
	} else {
		fmt.Println("Available Books:")
		fmt.Println("*************************************************")
		fmt.Printf("%-10v %-10v %-10v %10v \n", "ID", "Title", "Author", "Status")
		for _, book := range books {
			fmt.Printf("%-10v %-10v %-10v %10v \n", book.ID, book.Title, book.Author, book.Status)
		}
	}

	separetor()

}

func ListBorrowedBooksController(reader *bufio.Reader) {
	// This method returns the list of books borrowed by a specifc member
	memberID, err := getMemberID(reader)

	if err != nil {
		fmt.Println("Invalid Member ID:", err)
		return
	}

	books, err := services.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else if len(books) == 0 {
		fmt.Println("No books have been borrowed by the given member")
	} else {
		fmt.Printf("Member ID: %-15v", memberID)
		fmt.Println("Borrowed Books:")
		fmt.Println("*************************************************")
		fmt.Printf("%-10v %-10v %-10v %10v \n", "ID", "Title", "Author", "Status")
		for _, book := range books {
			fmt.Printf("%-10v %-10v %-10v %10v \n", book.ID, book.Title, book.Author, book.Status)
		}
	}

	separetor()

}
