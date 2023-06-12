package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/controller"
)

type authRouter struct {
	controller controller.BaseController
}

func NewAuthRoute(baseController controller.BaseController) BaseRoute {
	return &authRouter{controller: baseController}
}

func (u authRouter) RegisterRoutes(app *fiber.App) {
	app.Post("/auth/signup", u.controller.Add)
	app.Get("/auth/signup", u.controller.Add)
	app.Add("POST", "/auth/signin", u.controller.Get)
	app.Add("GET", "/auth/signin", u.controller.Get)
	app.Delete("/auth/signout", u.controller.Delete)
}
