package postgresdb

import (
	"book_store/internal/domain"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// BookRepositoryDb is interface implementation (adapter) for the secondary port interface BookRepository
type BookRepositoryPostgreSQL struct {
	client *sqlx.DB
}

var book domain.Book

func (b BookRepositoryPostgreSQL) GetBooks(ctx *gin.Context) ([]domain.Book, error) {
	var books []domain.Book

	sqlRequest := "select * from books_store"
	if err := b.client.Select(&books, sqlRequest); err != nil {
		return nil, err
	}

	return books, nil
}

func (b BookRepositoryPostgreSQL) GetBook(ctx *gin.Context, id int) (*domain.Book, error) {

	sqlRequest := "select * from books_store where id = $1"
	if err := b.client.Get(&book, sqlRequest, id); err != nil {
		return nil, err
	}

	return &book, nil
}

func (b BookRepositoryPostgreSQL) NewBook(ctx *gin.Context, req domain.Book) (int, error) {

	sqlRequest := "INSERT INTO books_store (Title, Authors, Year) VALUES ($1, $2, $3)"

	res, err := b.client.Exec(sqlRequest, req.Title, req.Authors, req.Year)
	if err != nil {
		return 0, err
	}

	rowsAdded, _ := res.RowsAffected()

	return int(rowsAdded), nil
}

func (b BookRepositoryPostgreSQL) DeleteBook(ctx *gin.Context, id int) (int, error) {

	sqlRequest := "DELETE FROM books_store where id = $1"

	res, err := b.client.Exec(sqlRequest, id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, _ := res.RowsAffected()
	return int(rowsDeleted), nil
}

func (b BookRepositoryPostgreSQL) UpdateBook(ctx *gin.Context, req domain.Book) (int, error) {

	result, err := b.client.Exec("UPDATE books_store SET Title=$1, Authors=$2, Year=$3 where id=$4 RETURNING id", req.Title, req.Authors, req.Year, req.ID)
	if err != nil {
		return 0, err
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
	db_user := os.Getenv("POSTGRES_USER")
	db_pswd := os.Getenv("POSTGRES_PASSWORD")
	db_address := os.Getenv("DB_ADDRESS")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("POSTGRES_DB")

	// checkEnvVars verifies if all env variables have been set
	checkEnvVars("POSTGRES_USER", "POSTGRES_PASSWORD", "DB_ADDRESS", "DB_PORT", "POSTGRES_DB")

	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", db_address, db_port, db_user, db_name, db_pswd)	
	client, err := sqlx.Open("postgres", dataSource)
	if err != nil || client == nil {
		log.Fatal("Error while opening DB: ", err.Error())
	}

	err = client.Ping()
	if err != nil {
		log.Fatalf("DBConnection error: %s", err.Error())
	}

	// Reading file with SQL instructions 
	res, err1 := ioutil.ReadFile("instructions.sql")
	if err1 != nil {
		log.Fatalf("Error while reading file with instructions: %v", err1.Error())
	}
	var schema = string(res)
	client.MustExec(schema)

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func checkEnvVars(s ...string) {
	for _, j := range s {
		if _, boolean := os.LookupEnv(j); !boolean {
			log.Panicf("Env variable %v is not set", j)
		}
	}
}
