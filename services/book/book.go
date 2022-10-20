package book

import (
	"learn_orm/dto"
	"learn_orm/models"
	"learn_orm/repository"
)

type BookService interface {
	GetAll() ([]dto.DTOBookRes, error)
	GetById(id uint) (dto.DTOBookRes, error)
	Create(payload dto.DTOBookReq) error
	Update(id uint, payload dto.DTOBookReq) error
	Delete(id uint) error
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(rep repository.BookRepository) *bookService {
	return &bookService{rep}
}

func (srv *bookService) GetAll() ([]dto.DTOBookRes, error) {
	dtoBooks := []dto.DTOBookRes{}

	books, err := srv.bookRepository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		dtoBook := dto.DTOBookRes{
			ID:          book.ID,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
			DeletedAt:   book.DeletedAt,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			Publisher:   book.Publisher,
		}

		dtoBooks = append(dtoBooks, dtoBook)
	}

	return dtoBooks, err
}

func (srv *bookService) GetById(id uint) (dto.DTOBookRes, error) {
	dtoBook := dto.DTOBookRes{}

	book, err := srv.bookRepository.GetById(id)
	if err != nil {
		return dtoBook, err
	}

	dtoBook.ID = book.ID
	dtoBook.CreatedAt = book.CreatedAt
	dtoBook.UpdatedAt = book.UpdatedAt
	dtoBook.DeletedAt = book.DeletedAt
	dtoBook.Title = book.Title
	dtoBook.Author = book.Author
	dtoBook.Description = book.Description
	dtoBook.Publisher = book.Publisher

	return dtoBook, err
}

func (srv *bookService) Create(payload dto.DTOBookReq) error {
	book := models.Book{}
	book.Title = payload.Title
	book.Author = payload.Author
	book.Description = payload.Description
	book.Publisher = payload.Publisher

	return srv.bookRepository.Create(book)
}

func (srv *bookService) Update(id uint, payload dto.DTOBookReq) error {

	book := models.Book{}
	book.Title = payload.Title
	book.Author = payload.Author
	book.Description = payload.Description
	book.Publisher = payload.Publisher

	return srv.bookRepository.Update(id, book)
}

func (srv *bookService) Delete(id uint) error {
	return srv.bookRepository.Delete(id)
}
