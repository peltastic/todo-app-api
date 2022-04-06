package routes

import (
	"todo/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("api/todo", controllers.CreateTodo)
	app.Post("api/logout", controllers.Logout)
	app.Delete("api/delete/:id", controllers.DeleteTodo)
	app.Put("api/update/:id", controllers.TodoCompleted)
}
