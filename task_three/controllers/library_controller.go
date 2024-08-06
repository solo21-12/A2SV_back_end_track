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

const BORROWED string = "Borrowed"
const AVAILABLE string = "Available"

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), err
}

func parseStr(input string) (int, error) {
	res, err := strconv.Atoi(input)

	return res, err

}

func generateNewBookID() int {
	return len(data.OurLibrary.Books) + 1
}

func generateNewMemberID() int {
	return len(data.OurLibrary.Members)
}

func getMemberAndBook(reader *bufio.Reader) (int, int, error) {
	bid, _ := getInput("Enter Book ID to borrow: ", reader)
	mid, _ := getInput("Enter Member ID: ", reader)

	bookID, bErr := parseStr(bid)
	memberID, mErr := parseStr(mid)

	if bErr != nil {
		fmt.Println("Invalid Book ID:", bErr)
		return -1, -1, bErr
	}

	if mErr != nil {
		fmt.Println("Invalid Member ID:", mErr)
		return -1, -1, mErr

	}

	return bookID, memberID, nil

}

func AddBook(reader *bufio.Reader) {
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

}

func AddMember(reader *bufio.Reader) {
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

}

func RemoveBook(reader *bufio.Reader) {
	id, _ := getInput("Enter Book ID to remove: ", reader)
	bookID, _ := strconv.Atoi(id)
	err := services.RemoveBook(bookID)

	if err != nil {
		fmt.Println("Error removing book:", err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func BorrowBook(reader *bufio.Reader) {

	bookID, memberID, cur_err := getMemberAndBook(reader)

	if cur_err != nil {
		fmt.Println("Error while getting information", cur_err)
	} else {
		err := services.BorrowBook(bookID, memberID)

		if err != nil {
			fmt.Println("Error borrowing book: ", err)
		} else {
			fmt.Println("Book borrowed successfully.")
		}
	}

}

func ReturnBook(reader *bufio.Reader) {
	bookID, memberID, cur_err := getMemberAndBook(reader)

	if cur_err != nil {
		fmt.Println("Error while getting information", cur_err)
	} else {
		err := services.ReturnBook(bookID, memberID)

		if err != nil {
			fmt.Println("Error returning book: ", err)
		} else {
			fmt.Println("Book returned successfully.")
		}
	}

}

func ListAvailableBooks() {
	books := services.ListAvailableBooks()
	fmt.Println("Available Books:")
	fmt.Println("*************************************************")
	fmt.Printf("%-10v %-10v %-10v %10v \n", "ID", "Title", "Author", "Status")

	for _, book := range books {
		fmt.Printf("%-10v %-10v %-10v %10v \n", book.ID, book.Title, book.Author, book.Status)
	}
}

func ListBorrowedBooks(reader *bufio.Reader) {
	mid, _ := getInput("Enter Member ID: ", reader)
	memberID, err := parseStr(mid)
	if err != nil {
		fmt.Println("Invalid Member ID:", err)
		return
	}

	books := services.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No books have been borrowed by the given member")
	} else {
		fmt.Println("Borrowed Books:")
		fmt.Println("*************************************************")
		fmt.Printf("%-10v %-10v %-10v %10v \n", "ID", "Title", "Author", "Status")
		for _, book := range books {
			fmt.Printf("%-10v %-10v %-10v %10v \n", book.ID, book.Title, book.Author, book.Status)
		}
	}
}
