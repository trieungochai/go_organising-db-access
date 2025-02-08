package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/trieungochai/go_organising-db-access/02.dependency-injection/models"
)

// Create a custom Env struct which holds a connection pool.
type Env struct {
	db *sql.DB
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

	// Initialise the connection pool.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of Env containing the connection pool.
	env := &Env{db: db}

	// Use env.booksIndex as the handler function for the /books route.
	http.HandleFunc("/books", env.booksIndex)

	log.Println("Server is running on port 3000...")
	http.ListenAndServe(":3000", nil)
}

// Define booksIndex as a method on Env.
func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	books, err := models.AllBooks(env.db)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, book := range books {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", book.Isbn, book.Title, book.Author, book.Price)
	}
}
