package utils

import (
	"go-bookman-app/services"
	"log"
	"time"
)

func InsertBooks(dbService services.BookDBService, cacheService services.BookCacheService) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		books, err := dbService.GetAllBooks()

		if err != nil {
			log.Printf("fetch books failed: %v\n", err)
		}

		cacheService.InsertBatchBook(books)
	}
}
