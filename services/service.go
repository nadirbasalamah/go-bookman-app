package services

import (
	"go-bookman-app/models"
	"reflect"

	"github.com/redis/go-redis/v9"
)

type BookService struct {
	dbService    BookDBService
	cacheService BookCacheService
}

func InitBookService(dbService BookDBService, cacheService BookCacheService) BookService {
	return BookService{
		dbService:    dbService,
		cacheService: cacheService,
	}
}

func (b *BookService) GetAllBooks() ([]models.Book, error) {
	books, err := b.cacheService.GetAllBooks()

	isErr := err != nil && err != redis.Nil

	if isErr {
		return nil, err
	}

	if len(books) > 0 {
		return books, nil
	}

	return b.dbService.GetAllBooks()
}

func (b *BookService) GetBookByID(id string) (models.Book, error) {
	book, err := b.cacheService.GetBookByID(id)

	isErr := err != nil && err != redis.Nil

	if isErr {
		return models.Book{}, err
	}

	if err != redis.Nil {
		return book, nil
	} else {
		book, err := b.dbService.GetBookByID(id)
		isEmpty := reflect.ValueOf(book).IsZero()

		if err != nil || isEmpty {
			return models.Book{}, err
		}

		b.cacheService.InsertBook(book)
	}

	book, err = b.cacheService.GetBookByID(id)

	isErr = err != nil && err != redis.Nil

	if isErr {
		return models.Book{}, err
	}

	return book, nil
}

func (b *BookService) InsertBook(request models.BookRequest) (models.Book, error) {
	newBook := models.Book{
		Title:       request.Title,
		Description: request.Description,
		Publisher:   request.Publisher,
	}

	book, err := b.dbService.InsertBook(newBook)

	if err != nil {
		return models.Book{}, err
	}

	return b.cacheService.InsertBook(book)
}

func (b *BookService) InsertBatchBook(request []models.BookRequest) error {
	newBooks := make([]models.Book, len(request))

	for idx, req := range request {
		newBooks[idx] = models.Book{
			Title:       req.Title,
			Description: req.Description,
			Publisher:   req.Publisher,
		}
	}

	return b.dbService.InsertBatchBook(newBooks)
}
