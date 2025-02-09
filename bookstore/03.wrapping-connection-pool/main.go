package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/trieungochai/go_organising-db-access/03.wrapping-connection-pool/models"
)

// This time make models.BookModel the dependency in Env.
type Env struct {
	books models.BookModel
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL env variable is not set")
	}

	// Init the connection pool as normal.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Initalise Env with a models.BookModel instance
	// (which in turn wraps the connection pool).
	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksIndex)
	log.Println("Server is running on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	// Execute the SQL query by calling the AllBooks() method.
	books, err := env.books.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, book := range books {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", book.Isbn, book.Title, book.Author, book.Price)
	}
}
