package models

type Library struct {
	Books   map[int]Book
	Members map[int]Member
}
