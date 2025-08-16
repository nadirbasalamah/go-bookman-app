package routes

import (
	"go-bookman-app/controllers"
	"go-bookman-app/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, bookService services.BookService) {
	bookController := controllers.InitBookController(bookService)

	bookRoutes := e.Group("/api/v1")

	bookRoutes.GET("/books", bookController.GetAllBooks)
	bookRoutes.GET("/books/:id", bookController.GetBookByID)
	bookRoutes.POST("/books", bookController.CreateBook)
	bookRoutes.POST("/books/batch", bookController.CreateBatchBook)
}
