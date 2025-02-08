package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// This time the global variable is unexported.
var db *sql.DB

// InitDB sets up setting up the connection pool global variable.
func Init(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	return db.Ping()
}

// The Book struct represents a book entity
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]Book, error) {
	// a. Querying the Database
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	// ensures that the result set is closed when the func exits, preventing resource leaks.
	defer rows.Close()

	// b. Iterating Over the Query Results
	// Init an empty slice bks to store the retrieved books.
	var books []Book

	// The rows.Next() function moves the cursor to the next row.
	// rows.Scan() extracts values from the current row into the provided variables.
	// This process continues until rows.Next() returns false, meaning there are no more rows.
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	// c. Handling Potential Errors
	// Checks if any error occurred during the iteration (rows.Next() loop).
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
