package main

import (
	"os"

	"github.com/bimaagung/fiber-crud/config"
	"github.com/bimaagung/fiber-crud/controllers"
	"github.com/gofiber/fiber/v2"
)


func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/list", controllers.AddList)
	v1.Get("/list", controllers.GetAllList)
	v1.Get("/list/:id", controllers.GetListById)
	v1.Put("/list/:id", controllers.UpdateList)
	v1.Delete("/list/:id", controllers.DeleteList)

	app.Get("/", func(c *fiber.Ctx) error{
		return c.SendString("Hello, world!")
	})

	app.Listen(":"+os.Getenv("PORT"))
}