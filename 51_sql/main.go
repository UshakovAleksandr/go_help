package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func deleteBook(db *sql.DB) {
	query := "delete from books where rating < 4"
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	fmt.Printf("deleted %d books, error=%v\n", count, err)
	// deleted 3 books, error=<nil>
}

func updateBook(db *sql.DB) {
	query := "update books set author = ? where author = ?"
	res, err := db.Exec(query, "J.R.R! Tolkien", "J.R.R. Tolkin")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	fmt.Printf("updated %d books, error=%v\n", count, err)
}

func insertBooks(db *sql.DB) {
	query := `
    insert into books(title, author, num_pages, rating)
    values (?, ?, ?, ?)
`

	data := [][]any{
		{"The Catcher in the Rye", "J.D. Salinger", 277, 3.8},
		{"The Fellowship of the Ring", "J.R.R. Tolkin", 398, 4.36},
		{"The Giver", "Lois Lowry", 208, 4.13},
		{"The Da Vinci Code", "Dan Brown", 489, 3.84},
		{"The Alchemist", "Paulo Coelho", 197, 3.86},
	}

	for _, vals := range data {
		res, err := db.Exec(query, vals...)
		if err != nil {
			panic(err)
		}

		bookID, err := res.LastInsertId()
		fmt.Printf("added new book: id=%d, error=%v\n", bookID, err)
	}
}

func createTable(db *sql.DB) {
	query := `
    drop table if exists books;
    create table if not exists books(
        id integer primary key,
        title text,
        author text,
        num_pages integer,
        rating real
    );
`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("✓ created books table")
}

func main() {
	// подключиться к БД
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("✓ connected to books db")
	// ✓ connected to books db

	createTable(db)
	//insertBooks(db)
	updateBook(db)
}
