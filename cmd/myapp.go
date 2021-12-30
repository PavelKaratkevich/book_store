package main

import (
	"book_store/internal/handler"
	"book_store/internal/service"
	postgresdb "book_store/internal/repositoryDB/postgresDB"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create DB client for PostgreSQL
	db := postgresdb.ConnectDB()

	// Wiring
	bookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(bookRepositoryDB)
	bh := handler.BookHandler{Service: bookService}

	// Creating router and defining handler
	g := gin.Default()
	g.GET("/books/", bh.GetAllBook)

	// Running the connection on a defined port
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}