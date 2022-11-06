package main

import (
	"github.com/bimaagung/fiber-crud/config"
	"github.com/bimaagung/fiber-crud/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.List{})
}