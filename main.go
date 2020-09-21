package main

import (
	"GoCrudRestExample/models"
	"github.com/gofiber/fiber"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/users", models.GetUsers)
	app.Post("/users", models.CreateUser)
	app.Get("/users/:id", models.GetUserById)
	app.Put("/users/:id", models.UpdateUser)
	app.Delete("/users/:id", models.DeleteUserById)

	log.Fatal(app.Listen(":3000"))
}
