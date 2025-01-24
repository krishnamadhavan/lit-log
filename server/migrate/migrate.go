package main

import (
	"github.com/krishnamadhavan/lit-log/initializers"
	"github.com/krishnamadhavan/lit-log/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Book{},
		&models.Author{},
		&models.Genre{},
		&models.Language{},
	)

	if err != nil {
		return
	}
}
