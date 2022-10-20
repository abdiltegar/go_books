package repository

import (
	"gorm.io/gorm"
	"learn_orm/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (models.User, error)
	GetByEmailPassword(email, password string) (models.User, error)
	Create(user models.User) error
	Update(id uint, user models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	users := []models.User{}

	if err := r.db.Model([]models.User{}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) GetById(id uint) (models.User, error) {
	user := models.User{}

	if err := r.db.Model(models.User{}).Where("ID = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetByEmailPassword(email, password string) (models.User, error) {
	user := models.User{}

	if err := r.db.Model(models.User{}).Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) Update(id uint, user models.User) error {
	return r.db.Model(models.User{}).Where("ID = ?", id).Updates(user).Error
}

func (r *userRepository) Delete(id uint) error {
	var user models.User
	return r.db.Delete(&user, id).Error
}
