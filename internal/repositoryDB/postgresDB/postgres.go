package postgresdb

import (
	"book_store/internal/domain"
	"book_store/internal/response"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// BookRepositoryDb is interface implementation (adapter) for the secondary port interface BookRepository
type BookRepositoryPostgreSQL struct {
	client *sqlx.DB
}

var books []domain.Book
var appError err.AppError
var book domain.Book

func (b BookRepositoryPostgreSQL) GetBooks() ([]domain.Book, *err.AppError) {
	sqlRequest := "select * from books_store"
	if err := b.client.Select(&books, sqlRequest); err != nil {
		appError.Message = "Unknown error"
		appError.Code = http.StatusInternalServerError
		return nil, &appError
	}
	return books, nil
}

func (b BookRepositoryPostgreSQL) GetBook(id int) (*domain.Book, *err.AppError) {
	sqlRequest := "select * from books_store where id = $1"
	if err := b.client.Get(&book, sqlRequest, id); err != nil {
		if err == sql.ErrNoRows {
			appError.Message = "ID not found"
			appError.Code = http.StatusNotFound
			return nil, &appError
		} else {
			appError.Message = "Unknown error"
			appError.Code = http.StatusInternalServerError
			return nil, &appError
		}
	}
	return &book, nil
}

// helper function
func NewBookRepositoryDb(client *sqlx.DB) BookRepositoryPostgreSQL {
	return BookRepositoryPostgreSQL{
		client,
	}
}

func ConnectDB() *sqlx.DB {
	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// get environment variables
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_address := os.Getenv("DB_ADDRESS")
	db_pswd := os.Getenv("DB_PSWD")

	dataSource := fmt.Sprintf("postgres://postgres:%s@%s:%s/%s?sslmode=disable", db_pswd, db_address, db_port, db_name)
	client, err := sqlx.Open("postgres", dataSource)
	if err != nil || client == nil {
		log.Fatal("Error while opening DB: ", err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
