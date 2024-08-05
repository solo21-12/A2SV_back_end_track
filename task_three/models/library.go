package models

type Library struct {
	books   map[int]Book
	members map[int]Member
}
