package controllers

import (
	"math/rand"
	"strconv"
	"todo/db"
	"todo/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateTodo(c *fiber.Ctx) error {
	var data models.Todo
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "can't create todo, not authenticated",
		})
	}
	todoid := rand.Intn(30000-10000) + 10000
	claims := token.Claims.(*jwt.RegisteredClaims)
	userid, err := strconv.Atoi(claims.Issuer)
	todo := models.Todo{
		UserID:      userid,
		Todo:        data.Todo,
		TodoID:      todoid,
		IsCompleted: false,
	}
	db.DB.Create(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	_, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "can't delete todo, not authenticated",
		})
	}
	db.DB.Where("todo_id = ?", c.Params("id")).Delete(&models.Todo{})
	return c.JSON(fiber.Map{
		"message": "deleted sucessfully",
	})
}

func TodoCompleted(c *fiber.Ctx) error {
	var todo models.Todo
	db.DB.Where("todo_id = ?", c.Params("id")).First(&todo)
	if todo.IsCompleted == false {
		db.DB.Model(&todo).Where("todo_id = ?", c.Params("id")).Update("is_completed", true)
	}
	db.DB.Model(&todo).Where("todo_id = ?", c.Params("id")).Update("is_completed", false)

	return nil
}
