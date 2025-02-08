package models

import "database/sql"

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// Update the AllBooks func so it accepts the connection pool as a param.
func AllBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
