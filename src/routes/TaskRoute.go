package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/controller"
)

type taskRouter struct {
	controller controller.Controller
}

func NewTaskRoute(baseController controller.Controller) Route {
	return &taskRouter{controller: baseController}
}

func (u taskRouter) RegisterRoutes(app *fiber.App) {
	app.Post("/tasks/add", u.controller.Add)
}

func (u *taskRouter) GetController() *controller.Controller {
	return &u.controller
}
