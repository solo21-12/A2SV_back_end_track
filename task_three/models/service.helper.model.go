package models

type ServiceHelper interface {
	GetBook(bookID int) (Book, bool)
	GetMember(memberID int) (Member, bool)
	CheckBookBorrowed(book Book, member Member) (int, bool)
	CheckBorrowed(book Book) bool
	CheckBookBorrowAllMemeber(bookID int) (int, int)
}
