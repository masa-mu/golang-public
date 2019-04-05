package message

import "github.com/jinzhu/gorm"

type Repository interface {
	GetMessage() (err error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) Repository {
	return &repository{DB}
}

func (r repository) GetMessage() (err error) {
	return nil
}
