package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title           string `json:"title"`
	LongDescription string `json:"long_description"`
	Format          string `json:"format"`
	Published       string `json:"published"`
	ISBN            int    `json:"isbn"`
	ASIN            int    `json:"asin"`
}

// add & link Language, Author, Genre
