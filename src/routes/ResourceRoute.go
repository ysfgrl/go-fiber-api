package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/controller"
)

type resourceRouter struct {
	controller controller.Controller
}

func NewResourceRoute(baseController controller.Controller) Route {
	return &resourceRouter{controller: baseController}
}

func (u resourceRouter) RegisterRoutes(app *fiber.App) {
	app.Get("/resources/list", u.controller.List)
	app.Post("/resources/add", u.controller.Add)
	app.Get("/resources/get/:id", u.controller.Get)
	app.Post("/resources/upload", u.controller.SetFile)
}

func (u *resourceRouter) GetController() *controller.Controller {
	return &u.controller
}
