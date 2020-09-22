package models

import (
	"GoCrudRestExample/db"
	"encoding/json"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

//var users = []User{
//	{
//		ID:       1,
//		Name:     "Joan",
//		LastName: "Cabezas",
//		Email:    "joan@concha.abs.com",
//		Age:      19,
//	},
//}

type User struct {
	gorm.Model
	ID       int
	Name     string
	LastName string
	Email    string
	Age      int
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	db.DBConn.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		panic(err)
	}
	db.DBConn.Create(&user)
	c.Status(http.StatusCreated)
	return c.JSON(user)
}

func GetUserById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("")
	}
	var user User
	db.DBConn.Find(&user, id)
	if user.ID != 0 {
		return c.JSON(user)
	}

	c.Status(http.StatusNotFound)
	return c.SendString("User not found")
}

func DeleteUserById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("")
	}
	var user User
	db.DBConn.Find(&user, id)
	if user.ID != 0 {
		db.DBConn.Delete(user)
		return c.SendString("User deleted")
	}
	c.Status(http.StatusNotFound)
	return c.SendString("User not found")
}

func UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString("")
	}
	var newUser User
	err = json.Unmarshal(c.Body(), &newUser)
	if err != nil {
		panic(err)
	}

	var user User
	db.DBConn.Find(&user, id)
	if user.ID != 0 {
		user.Name = newUser.Name
		user.LastName = newUser.LastName
		user.Email = newUser.Email
		user.Age = newUser.Age
		db.DBConn.Save(user)
		return c.JSON(user)
	}
	c.Status(http.StatusNotFound)
	return c.SendString("User not found")
}
