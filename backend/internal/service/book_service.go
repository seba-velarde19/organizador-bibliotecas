package service

import (
	"biblioteca/internal/models"
	"biblioteca/internal/store"
	"errors"
)

type Service struct {
	store store.Store
}

func NewBookService(store store.Store) *Service {
	return &Service{store: store}
}

func (s *Service) GetAllBooks() ([]*models.Book, error) {
	return s.store.GetAll()
}

func (s *Service) GetBook(id int) (*models.Book, error) {
	return s.store.GetByID(id)
}

func (s *Service) GetBookByTitle(title string) (*models.Book, error) {
	return s.store.GetByTitle(title)
}

func (s *Service) AddBook(book models.Book) (*models.Book, error) {
	if book.Title == "" {
		return nil, errors.New("book title is empty")
	}
	if book.Author == "" {
		return nil, errors.New("book author is empty")
	}
	return s.store.AddBook(&book)
}

func (s *Service) UpdateBook(id int, book models.Book) (*models.Book, error) {
	if book.Title == "" {
		return nil, errors.New("book title is empty")
	}
	if book.Author == "" {
		return nil, errors.New("book author is empty")
	}
	return s.store.UpdateBook(id, &book)
}

func (s *Service) DeleteBook(id int) error {
	return s.store.DeleteBook(id)
}
