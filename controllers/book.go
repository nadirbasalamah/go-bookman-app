package controllers

import (
	"go-bookman-app/models"
	"go-bookman-app/services"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	service services.BookService
}

func InitBookController(service services.BookService) BookController {
	return BookController{
		service: service,
	}
}

func (bc *BookController) GetAllBooks(c echo.Context) error {
	books, err := bc.service.GetAllBooks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "failed to fetch books",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Book]{
		Message: "all books",
		Data:    books,
	})
}

func (bc *BookController) GetBookByID(c echo.Context) error {
	bookID := c.Param("id")

	book, err := bc.service.GetBookByID(bookID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[any]{
			Message: "failed to fetch book",
		})
	}

	if reflect.ValueOf(book).IsZero() {
		return c.JSON(http.StatusNotFound, models.Response[any]{
			Message: "book not found",
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

	book, err := bc.service.InsertBook(request)

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

	err := bc.service.InsertBatchBook(request)

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
