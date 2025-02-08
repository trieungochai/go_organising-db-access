package models

import "database/sql"

// A global variable DB of type *sql.DB is declared to hold the db connection pool.
// This allows different parts of the app to use the same db connection
// without having to create new ones.
var DB *sql.DB

// The Book struct represents a book entity
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]Book, error) {
	// a. Querying the Database
	rows, err := DB.Query("SELECT * FROM books")
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
