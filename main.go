package main

import (
	"GoCrudRestExample/db"
	"GoCrudRestExample/models"
	"fmt"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func initDatabase() {
	var err error
	dsn := "postgres://plwrwgsqbtghhm:857d60c3c0d54656bc25bf81f3d6edc856b1ed9e80f32fed4ad2ab049df5395b@ec2-54-236-169-55.compute-1.amazonaws.com:5432/d7battfi3i1v7f"
	db.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.DBConn.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database migrated")
}

func main() {
	initDatabase()
	app := fiber.New()

	app.Get("/users", models.GetUsers)
	app.Post("/users", models.CreateUser)
	app.Get("/users/:id", models.GetUserById)
	app.Put("/users/:id", models.UpdateUser)
	app.Delete("/users/:id", models.DeleteUserById)

	log.Fatal(app.Listen(":3000"))
}
