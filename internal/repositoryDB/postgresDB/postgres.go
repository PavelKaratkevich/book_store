package postgresdb

import (
	"book_store/internal/domain"
	err "book_store/internal/response"
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

var appError err.AppError
var book domain.Book

func (b BookRepositoryPostgreSQL) GetBooks() ([]domain.Book, *err.AppError) {
	var books []domain.Book

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

func (b BookRepositoryPostgreSQL) NewBook(req domain.Book) (int, *err.AppError) {

	sqlRequest := "INSERT INTO books_store (Title, Authors, Year) VALUES ($1, $2, $3)"

	res, err := b.client.Exec(sqlRequest, req.Title, req.Authors, req.Year)
	if err != nil {
		log.Printf("Error while inserting book to DB: %v", err.Error())
		appError.Message = "Error while making a book record"
		appError.Code = http.StatusInternalServerError
		return 0, &appError
	}

	rowsAdded, _ := res.RowsAffected()

	return int(rowsAdded), nil
}

func (b BookRepositoryPostgreSQL) DeleteBook(id int) (int, *err.AppError) {

	sqlRequest := "DELETE FROM books_store where id = $1"

	res, err := b.client.Exec(sqlRequest, id)
	if err != nil {
		log.Printf("Error while deleting book from DB: %v", err.Error())
		appError.Message = fmt.Sprintf("Error while deleting book with %v from DB", id)
		appError.Code = http.StatusInternalServerError
		return 0, &appError
	}

	rowsDeleted, _ := res.RowsAffected()

	return int(rowsDeleted), nil
}

func (b BookRepositoryPostgreSQL) UpdateBook(req domain.Book) (int, *err.AppError) {

	result, err := b.client.Exec("UPDATE books_store SET Title=$1, Authors=$2, Year=$3 where id=$4 RETURNING id", req.Title, req.Authors, req.Year, req.ID)

	if err != nil {
		log.Printf("Error while updating book from DB: %v", err.Error())
		appError.Message = fmt.Sprintf("Error while updating book with %v from DB", req.ID)
		appError.Code = http.StatusInternalServerError
		return 0, &appError
	}

	rowsUpdated, _ := result.RowsAffected()

	return int(rowsUpdated), nil
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
