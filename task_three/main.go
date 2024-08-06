package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/task_three/controllers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

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
			controllers.AddBookController(reader)
		case 2:
			controllers.RemoveBookController(reader)
		case 3:
			controllers.BorrowBookController(reader)
		case 4:
			controllers.ReturnBookController(reader)
		case 5:
			controllers.ListAvailableBooksController()
		case 6:
			controllers.ListBorrowedBooksController(reader)
		case 7:
			controllers.AddMemberController(reader)
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
