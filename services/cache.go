package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bookman-app/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type BookCacheService struct {
	rdb *redis.Client
}

func InitBookCacheService(rdb *redis.Client) BookCacheService {
	return BookCacheService{
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

func (b *BookCacheService) InsertBook(book models.Book) (models.Book, error) {
	data, err := json.Marshal(&book)

	if err != nil {
		return models.Book{}, err
	}

	expire := 30 * time.Second

	bookID := fmt.Sprintf("book:%d", book.ID)

	res := b.rdb.SetNX(context.TODO(), bookID, data, expire)

	if res.Err() != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (b *BookCacheService) InsertBatchBook(books []models.Book) error {
	data, err := json.Marshal(&books)

	if err != nil {
		return err
	}

	expire := 30 * time.Second

	res := b.rdb.SetNX(context.TODO(), "books", data, expire)

	if res.Err() != nil {
		return err
	}

	return nil
}
