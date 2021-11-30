package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syncodd/simple-todoapp-fiber-go/controller"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controller.GetTodos)
	route.Post("", controller.CreateTodo)
	route.Put("/:id", controller.UpdateTodo)
	route.Delete("/:id", controller.DeleteTodo)
	route.Get("/:id", controller.GetTodo)
}
