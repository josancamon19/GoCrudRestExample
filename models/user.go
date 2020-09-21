package models

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"net/http"
	"strconv"
)

var users = []User{
	{
		ID:       1,
		Name:     "Joan",
		LastName: "Cabezas",
		Email:    "joan@concha.abs.com",
		Age:      19,
	},
}

type User struct {
	ID       int
	Name     string
	LastName string
	Email    string
	Age      int
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		panic(err)
	}
	users = append(users, user)
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
	for _, user := range users {
		if user.ID == id {
			return c.JSON(user)
		}
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
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.Status(http.StatusNoContent)
			return c.SendString("Deleted")
		}
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
	for i, user := range users {
		if user.ID == id {
			users[i] = newUser
			c.Status(http.StatusOK)
			return c.SendString("Updated")
		}
	}
	c.Status(http.StatusNotFound)
	return c.SendString("User not found")
}
