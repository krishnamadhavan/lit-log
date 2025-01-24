package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model

	Name  string `json:"name"`
	Books []Book `gorm:"foreignKey:LanguageID" json:"books"`
}
