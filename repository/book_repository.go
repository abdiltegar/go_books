package repository

import (
	"gorm.io/gorm"
	"learn_orm/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id uint) (models.Book, error)
	Create(user models.Book) error
	Update(id uint, user models.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	books := []models.Book{}

	if err := r.db.Model([]models.Book{}).Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func (r *bookRepository) GetById(id uint) (models.Book, error) {
	book := models.Book{}

	if err := r.db.Model(models.Book{}).Where("ID = ?", id).First(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Create(book models.Book) error {
	return r.db.Create(&book).Error
}

func (r *bookRepository) Update(id uint, book models.Book) error {
	return r.db.Model(models.Book{}).Where("ID = ?", id).Updates(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	var book models.Book
	return r.db.Delete(&book, id).Error
}
