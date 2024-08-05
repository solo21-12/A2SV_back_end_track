package services

import (
	"example.com/task_three/data"
	"example.com/task_three/models"
	"fmt"
)

func AddBook(book models.Book) {

	_, value := data.OurLibrary.Books[book.ID]

	if value {
		fmt.Println("The book with the given inforamtion already exists")
	} else {
		data.OurLibrary.Books[book.ID] = book
	}

}
