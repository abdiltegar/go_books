package mock

import (
	"errors"
	"gorm.io/gorm"
	"learn_orm/dto"
	"time"
)

type BookMock struct {
}

func NewBookService() *BookMock {
	return &BookMock{}
}

func (srv *BookMock) GetAll() ([]dto.DTOBookRes, error) {
	res := []dto.DTOBookRes{
		{
			ID:          1,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Title:       "Learn HTML 5",
			Description: "Lorem Ipsum",
			Author:      "Sebastian",
			Publisher:   "Gramedia",
		},
		{
			ID:          2,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   gorm.DeletedAt{},
			Title:       "Learn CSS",
			Description: "Lorem Ipsum",
			Author:      "Sebastian",
			Publisher:   "Gramedia",
		},
	}

	return res, nil
}

func (srv *BookMock) GetById(id uint) (dto.DTOBookRes, error) {
	res := dto.DTOBookRes{}

	if !(id == 1) {
		return res, errors.New("data not found")
	}

	return res, nil
}

func (srv *BookMock) Create(payload dto.DTOBookReq) error {
	return nil
}

func (srv *BookMock) Update(id uint, payload dto.DTOBookReq) error {
	return nil
}

func (srv *BookMock) Delete(id uint) error {
	return nil
}
