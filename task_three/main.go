package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/task_three/controllers"
	"example.com/task_three/data"
	"example.com/task_three/services"
	"example.com/task_three/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	serviceHelper := utils.NewServiceHelper(data.OurLibrary)
	controllerHelper := utils.NewControllerHelper(data.OurLibrary)
	bookService := services.NewBookService(data.OurLibrary, serviceHelper)
	bookController := controllers.NewBookController(bookService, controllerHelper)
	memberService := services.NewMemberService(data.OurLibrary, serviceHelper)
	memberController := controllers.NewMemberController(memberService, controllerHelper)

	// libraryCo

	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			bookController.AddBookController(reader)
		case 2:
			bookController.RemoveBookController(reader)
		case 3:
			bookController.BorrowBookController(reader)
		case 4:
			bookController.ReturnBookController(reader)
		case 5:
			bookController.ListAvailableBooksController()
		case 6:
			bookController.ListBorrowedBooksController(reader)
		case 7:
			memberController.AddMemberController(reader)
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
