package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bookman-app/models"

	"github.com/redis/go-redis/v9"
)

type BookCacheService struct {
	rdb *redis.Client
}

func InitBookCacheService(rdb *redis.Client) BookService {
	return &BookCacheService{
		rdb: rdb,
	}
}

func (b *BookCacheService) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	val, err := b.rdb.Get(context.TODO(), "books").Result()

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(val), &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookCacheService) GetBookByID(id string) (models.Book, error) {
	var book models.Book

	bookID := fmt.Sprintf("book:%s", id)

	val, err := b.rdb.Get(context.TODO(), bookID).Result()

	if err != nil {
		return models.Book{}, err
	}

	if err := json.Unmarshal([]byte(val), &book); err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b *BookCacheService) CreateBook(book models.Book) (models.Book, error) {

}

func (b *BookCacheService) CreateBatchBook(books []models.Book) error {

}
