package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/trieungochai/go_organising-db-access/01.using-a-global-var/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL env variable is not set")
	}

	// Use the InitDB function to initialise the global variable.
	err = models.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", booksIndex)

	log.Println("Server is running on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	books, err := models.AllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, book := range books {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", book.Isbn, book.Title, book.Author, book.Price)
	}
}
