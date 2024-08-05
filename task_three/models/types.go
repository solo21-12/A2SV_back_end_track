package taskthree

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

type Library struct {
	books   map[int]Book
	members map[int]Member
}

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}
