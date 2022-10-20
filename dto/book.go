package dto

import (
	"gorm.io/gorm"
	"time"
)

type DTOBookRes struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `json:"title" form:"title"`
	Description string         `json:"description" form:"description"`
	Author      string         `json:"author" form:"author"`
	Publisher   string         `json:"publisher" form:"publisher"`
}

type DTOBookReq struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Author      string `json:"author" form:"author"`
	Publisher   string `json:"publisher" form:"publisher"`
}
