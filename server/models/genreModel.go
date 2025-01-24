package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model

	Name  string `json:"name"`
	About string `json:"about"`
	Books []Book `gorm:"many2many:book_genres;" json:"books"`
}
