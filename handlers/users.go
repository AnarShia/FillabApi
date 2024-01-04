package handlers

import (
	"fmt"

	"github.com/AnarShia/FillabApi/database"
	"github.com/AnarShia/FillabApi/models"
	"github.com/gofiber/fiber/v2"
)

func ListUsers(c *fiber.Ctx) error {
	facts := []models.User{}
	database.Db.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateUser(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.Db.Db.Create(&user)
	return c.Status(200).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	database.Db.Db.Find(&user, id)
	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	database.Db.Db.Find(&user, id)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.Db.Db.Save(&user)
	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	database.Db.Db.First(&user, id)
	fmt.Println(user)
	database.Db.Db.Delete(&user)
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully deleted user",
	})
}
