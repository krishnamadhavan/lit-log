package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title      string   `json:"title"`
	About      string   `json:"about"`
	Format     string   `json:"format"`
	Published  string   `json:"published"`
	ISBN       int      `json:"isbn"`
	ASIN       int      `json:"asin"`
	LanguageID int      `gorm:"index" json:"language_id"`
	Language   Language `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"language"`
	AuthorID   int      `gorm:"index" json:"author_id"`
	Author     Author   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
	Genres     []Genre  `gorm:"many2many:book_genres;" json:"genres"`
}
