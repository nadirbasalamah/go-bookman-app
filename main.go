package main

import (
	"go-bookman-app/routes"
	"go-bookman-app/services"
	"go-bookman-app/storage"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := storage.ConnectDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v\n", err)
	}

	rdb, err2 := storage.ConnectRedis()
	if err2 != nil {
		log.Fatalf("failed to connect redis: %v\n", err)
	}

	dbService := services.InitBookDBService(db)
	cacheService := services.InitBookCacheService(rdb)

	e := echo.New()

	routes.InitRoutes(e, dbService, cacheService)

	log.Println("all storage connected")

	e.Logger.Fatal(e.Start(":1323"))
}
