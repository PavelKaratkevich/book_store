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
	postgresBookRepositoryDB := postgresdb.NewBookRepositoryDb(db)
	bookService := service.NewBookService(postgresBookRepositoryDB)
	bh := handler.BookHandler{Service: bookService}

	// Creating router and defining routes and handlers
	g := gin.Default()
	g.GET("/books/", bh.GetAllBook)
	g.GET("/books/:id/", bh.GetBookbyIdNumber)

	// Running the connection on a defined port
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}