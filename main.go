package main

import (
	"go-bookman-app/routes"
	"go-bookman-app/services"
	"go-bookman-app/storage"
	"go-bookman-app/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var err error

	db, err := storage.ConnectDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v\n", err)
	}

	rdb, err := storage.ConnectRedis()
	if err != nil {
		log.Fatalf("failed to connect redis: %v\n", err)
	}

	dbService := services.InitBookDBService(db)
	cacheService := services.InitBookCacheService(rdb)
	service := services.InitBookService(
		dbService, cacheService,
	)

	go utils.InsertBooks(dbService, cacheService)

	e := echo.New()
	e.Use(middleware.Logger())

	routes.InitRoutes(e, service)

	log.Println("all storage connected")

	e.Logger.Fatal(e.Start(":1323"))
}
