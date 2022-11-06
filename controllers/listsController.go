package controllers

import (
	"github.com/bimaagung/fiber-crud/config"
	"github.com/bimaagung/fiber-crud/models"
	"github.com/gofiber/fiber/v2"
)

func AddList(c *fiber.Ctx) error {
	var body struct {
		Title  string
		Author string
		Body   string
	}

	c.BodyParser(&body)
	
	post := models.List{Title: body.Title,Author: body.Author, Body: body.Body}
	result := config.DB.Create(&post) 

	if result.Error != nil {
		 return c.Status(fiber.StatusBadRequest).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func GetAllList(c *fiber.Ctx) error {

	var list []models.List
	config.DB.Find(&list) 

	return c.Status(fiber.StatusOK).JSON(list)
}

func GetListById(c *fiber.Ctx) error {

	var list []models.List

	id := c.Params("id")

	config.DB.First(&list, id)

	return c.Status(fiber.StatusOK).JSON(list)
}

func UpdateList(c *fiber.Ctx) error {
	
	id := c.Params("id")
	
	var body struct {
		Body string
		Author string
		Title string
	}

	c.BodyParser(&body)

	var list []models.List
	config.DB.First(&list, id)
	config.DB.Model(&list).Updates(models.List{Title: body.Title, Author: body.Author,Body: body.Body})

	return c.Status(fiber.StatusOK).JSON(list)
}

func DeleteList(c *fiber.Ctx) error {
	
	id := c.Params("id")
	config.DB.Delete(&models.List{}, id)

	return c.Status(fiber.StatusOK).Context().Err()
}
