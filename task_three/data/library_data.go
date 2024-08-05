package data

import "example.com/task_three/models"

var OurLibrary = models.Library{
	Books:        make(map[int]models.Book),
	Members: make(map[int]models.Member),
}
