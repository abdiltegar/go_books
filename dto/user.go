package dto

import (
	"gorm.io/gorm"
	"time"
)

type DTOUserReq struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type DTOUserRes struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" form:"name"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password" form:"password"`
}

type DTOLoginRes struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
