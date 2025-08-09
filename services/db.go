package services

import (
	"go-bookman-app/models"

	"gorm.io/gorm"
)

type BookDBService struct {
	db *gorm.DB
}

func InitBookDBService(db *gorm.DB) BookService {
	return &BookDBService{
		db: db,
	}
}

func (b *BookDBService) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookDBService) GetBookByID(id string) (models.Book, error) {
	var book models.Book

	if err := b.db.Find(&book, id).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b *BookDBService) CreateBook(book models.Book) (models.Book, error) {
	if err := b.db.Create(&book).Error; err != nil {
		return models.Book{}, err
	}

	if err := b.db.First(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b *BookDBService) CreateBatchBook(books []models.Book) error {
	if err := b.db.Create(&books).Error; err != nil {
		return err
	}

	return nil
}
