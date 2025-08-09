package main

import (
	"go-bookman-app/routes"
	"go-bookman-app/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	storage.ConnectDB()
	storage.ConnectRedis()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
