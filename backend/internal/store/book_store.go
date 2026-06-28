package store

import (
	"biblioteca/internal/models"
	"database/sql"
)

type Store interface {
	GetAll() ([]*models.Book, error)
	GetByID(id int) (*models.Book, error)
	GetByTitle(title string) (*models.Book, error)
	AddBook(book *models.Book) (*models.Book, error)
	UpdateBook(id int, book *models.Book) (*models.Book, error)
	DeleteBook(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*models.Book, error) {
	query := "SELECT id,title,author FROM books"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []*models.Book
	for rows.Next() {
		var book *models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *store) GetByID(id int) (*models.Book, error) {
	var book *models.Book
	query := "SELECT id,title,author FROM books WHERE id=?"
	err := s.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *store) GetByTitle(title string) (*models.Book, error) {
	var book *models.Book
	query := "SELECT id,title,author FROM books WHERE title=?"
	err := s.db.QueryRow(query, title).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *store) AddBook(book *models.Book) (*models.Book, error) {
	query := "INSERT INTO books (title,author) VALUES (?,?)"
	result, err := s.db.Exec(query, book.Title, book.Author)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	book.ID = int(id)
	return book, nil
}

func (s *store) UpdateBook(id int, book *models.Book) (*models.Book, error) {
	query := "UPDATE books SET title=?, author=? WHERE id=?"
	_, err := s.db.Exec(query, book.Title, book.Author, id)
	if err != nil {
		return nil, err
	}
	book.ID = int(id)
	return book, nil
}

func (s *store) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id=?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
