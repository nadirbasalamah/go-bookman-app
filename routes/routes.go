package routes

import (
	"go-bookman-app/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	bookRoutes := e.Group("/api/v1")

	bookRoutes.GET("/books", controllers.GetAllBooks)
	bookRoutes.GET("/books/:id", controllers.GetBookByID)
	bookRoutes.POST("/books", controllers.CreateBook)
	bookRoutes.POST("/books/batch", controllers.CreateBatchBook)
}
