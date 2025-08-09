package services

import "go-bookman-app/models"

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id string) (models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	CreateBatchBook(books []models.Book) error
}
