package db

import (
	"book-library-service/internal/book-library-service/book"
	"book-library-service/pkg/logging"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type mysqldb struct {
	db     *sql.DB
	logger *logging.Logger
}

func (d *mysqldb) Create(ctx context.Context, b book.Book) (string, error) {
	result, err := d.db.Exec("INSERT INTO test.books (_id, name, type, content, author, year, description )"+
		" VALUES (?,?, ?, ?, ?, ?, ?)", uuid.New(), b.Name, b.Type, b.Content, b.Author, b.Year, b.Description)
	if err != nil {
		return "", fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "0", fmt.Errorf("addAlbum: %v", err)
	}
	return string(id), nil
}

func (d *mysqldb) FindAll(ctx context.Context) (books []book.Book, err error) {
	rows, err := d.db.Query("SELECT * FROM test.books")
	if err != nil {
		return nil, fmt.Errorf("failed to find all book. error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var b book.Book
		if err := rows.Scan(&b.ID, &b.Name, &b.Type, &b.Content, &b.Author, &b.Year, &b.Description); err != nil {
			return nil, fmt.Errorf("failed to find all book. error: %v", err)
		}
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to find all book. error: %v", err)
	}

	return books, nil
}

func (d *mysqldb) FindOne(ctx context.Context, id string) (b book.Book, err error) {
	row := d.db.QueryRow("SELECT * FROM test.books WHERE _id = ?", id)
	if err := row.Scan(&b.ID, &b.Name, &b.Type, &b.Content, &b.Author, &b.Year, &b.Description); err != nil {
		if err == sql.ErrNoRows {
			return b, fmt.Errorf("failed to find one book by id: %s: due to error %v", id, err)
		}
		return b, fmt.Errorf("failed to find one book by id: %s: due to error %v", id, err)
	}
	return b, nil
}

func (d *mysqldb) Update(ctx context.Context, book book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (d *mysqldb) Delete(ctx context.Context, id string) error {
	result, err := d.db.Exec("DELETE FROM test.books WHERE _id =?", id)
	if err != nil {
		return fmt.Errorf("failed to delete: %v", err)
	}
	fmt.Println(result.LastInsertId())
	return nil
}

func NewRepositoryM(database *sql.DB, logger *logging.Logger) book.Repository {

	return &mysqldb{
		db:     database,
		logger: logger,
	}

}
