package models

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	gorm.Model

	Name     string    `json:"name"`
	Born     time.Time `json:"born"`
	Died     time.Time `json:"died"`
	Website  string    `json:"website"`
	About    string    `json:"about"`
	PhotoUrl string    `json:"photo_url"`
	Books    []Book    `gorm:"foreignKey:AuthorID" json:"books"`
}

// TODO: add & link Genre
