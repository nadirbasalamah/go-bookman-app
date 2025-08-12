package services

import "go-bookman-app/models"

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id string) (models.Book, error)
	InsertBook(book models.Book) (models.Book, error)
	InsertBatchBook(books []models.Book) error
}
