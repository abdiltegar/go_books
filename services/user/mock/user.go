package mock

import (
	"errors"
	"gorm.io/gorm"
	"learn_orm/dto"
	"learn_orm/middlewares"
	"time"
)

type UserMock struct {
}

func NewUserMock() *UserMock {
	return &UserMock{}
}

func (srv *UserMock) GetAll() ([]dto.DTOUserRes, error) {
	res := []dto.DTOUserRes{
		{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
			Name:      "Admin",
			Email:     "admin@mail.com",
			Password:  "admin123",
		},
		{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
			Name:      "Fulan",
			Email:     "fulan@mail.com",
			Password:  "fulan123",
		},
	}

	return res, nil
}

func (srv *UserMock) GetById(id uint) (dto.DTOUserRes, error) {
	res := dto.DTOUserRes{}

	if !(id == 1) {
		return res, errors.New("data not found")
	}

	return res, nil
}

func (srv *UserMock) Create(payload dto.DTOUserReq) error {
	return nil
}

func (srv *UserMock) Update(id uint, payload dto.DTOUserReq) error {
	return nil
}

func (srv *UserMock) Delete(id uint) error {
	return nil
}

func (srv *UserMock) Login(payload dto.DTOUserReq) (dto.DTOLoginRes, error) {
	res := dto.DTOLoginRes{}

	if !(payload.Email == "admin@mail.com" && payload.Password == "admin123") {
		return res, errors.New("data not found")
	}

	token, _ := middlewares.CreateToken(1, payload.Email)

	res.ID = 1
	res.Name = "Admin"
	res.Email = payload.Email
	res.Token = token

	return res, nil
}
