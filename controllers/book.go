package controllers

import (
	"go-bookman-app/models"
	"go-bookman-app/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type BookController struct {
	dbService    services.BookDBService
	cacheService services.BookCacheService
}

func InitBookController(
	dbService services.BookDBService, cacheService services.BookCacheService) BookController {
	return BookController{
		dbService:    dbService,
		cacheService: cacheService,
	}
}

func (bc *BookController) GetAllBooks(c echo.Context) error {
	books, err := bc.cacheService.GetAllBooks()

	isErr := err != nil && err != redis.Nil

	if isErr {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "failed to fetch books",
		})
	}

	if len(books) > 0 {
		return c.JSON(http.StatusOK, models.Response[[]models.Book]{
			Message: "books from cache",
			Data:    books,
		})
	}

	books, err = bc.dbService.GetAllBooks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "failed to fetch books",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Book]{
		Message: "books from db",
		Data:    books,
	})
}

func (bc *BookController) GetBookByID(c echo.Context) error {
	bookID := c.Param("id")

	book, err := bc.cacheService.GetBookByID(bookID)

	isErr := err != nil && err != redis.Nil

	if isErr {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "failed to fetch book",
		})
	}

	if err != redis.Nil {
		return c.JSON(http.StatusOK, models.Response[models.Book]{
			Message: "book found",
			Data:    book,
		})
	} else {
		book, err := bc.dbService.GetBookByID(bookID)

		if err != nil {
			return c.JSON(http.StatusNotFound, models.Response[any]{
				Message: "book not found",
			})
		}

		bc.cacheService.InsertBook(book)
	}

	book, err = bc.cacheService.GetBookByID(bookID)

	isErr = err != nil && err != redis.Nil

	if isErr {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "failed to fetch book",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Book]{
		Message: "book found",
		Data:    book,
	})
}

func (bc *BookController) CreateBook(c echo.Context) error {
	var request models.BookRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Message: "failed to parse request",
		})
	}

	newBook := models.Book{
		Title:       request.Title,
		Description: request.Description,
		Publisher:   request.Publisher,
	}

	book, err := bc.dbService.InsertBook(newBook)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "insert book failed",
		})
	}

	book, err = bc.cacheService.InsertBook(book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "insert book failed",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Book]{
		Message: "book inserted",
		Data:    book,
	})
}

func (bc *BookController) CreateBatchBook(c echo.Context) error {
	var request []models.BookRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[any]{
			Message: "failed to parse request",
		})
	}

	newBooks := make([]models.Book, len(request))

	for idx, req := range request {
		newBooks[idx] = models.Book{
			Title:       req.Title,
			Description: req.Description,
			Publisher:   req.Publisher,
		}
	}

	err := bc.dbService.InsertBatchBook(newBooks)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "insert book failed",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[bool]{
		Message: "book inserted",
		Data:    true,
	})
}
