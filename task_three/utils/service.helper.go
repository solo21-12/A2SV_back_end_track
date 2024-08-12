package utils

import "example.com/task_three/models"

const (
	BORROWED  = "Borrowed"
	AVAILABLE = "Available"
)

type serviceHelper struct {
	data models.Library
}

func NewServiceHelper(data models.Library) models.ServiceHelper {
	return &serviceHelper{
		data: data,
	}
}

func (h *serviceHelper) GetBook(bookID int) (models.Book, bool) {
	book, exists := h.data.Books[bookID]
	return book, exists
}

func (h *serviceHelper) GetMember(memberID int) (models.Member, bool) {
	member, exist := h.data.Members[memberID]

	return member, exist
}

func (h *serviceHelper) CheckBookBorrowed(book models.Book, member models.Member) (int, bool) {

	for i, cur_book := range h.data.Members[member.ID].BorrowedBooks {
		if cur_book == book {
			return i, true
		}
	}

	return -1, false
}

func (h *serviceHelper) CheckBorrowed(book models.Book) bool {
	for _, cur_book := range h.data.Books {
		if cur_book == book && cur_book.Status == BORROWED {
			return true
		}
	}

	return false
}

func (h *serviceHelper) CheckBookBorrowAllMemeber(bookID int) (int, int) {
	for key, member := range h.data.Members {
		for i, cur_book := range member.BorrowedBooks {
			if cur_book.ID == bookID {
				return key, i
			}
		}
	}

	return -1, -1
}
