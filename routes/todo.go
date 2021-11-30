package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syncodd/simple-todoapp-fiber-go/controller"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controller.GetTodos)
	route.Post("", controller.CreateTodo)
	route.Put("/:id", controller.UpdateTodo) // curl -i -X PUT -H "Content-Type: application/json" -d "{""title"":""saamet""}" 127.0.0.1:8000/api/todos/1
	route.Delete("/:id", controller.DeleteTodo)
	route.Get("/:id", controller.GetTodo)
}
