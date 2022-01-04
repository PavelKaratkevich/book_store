package main

import (
	"book_store/internal/injection"
	"log"
)

func main() {
	// Start application
	router := injection.StartApp()
	
	// Running the connection on a defined port
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}