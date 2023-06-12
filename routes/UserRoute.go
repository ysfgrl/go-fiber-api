package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/controller"
)

type userRouter struct {
	controller controller.BaseController
}

func NewUserRoute(baseController controller.BaseController) BaseRoute {
	return &userRouter{controller: baseController}
}

func (u userRouter) RegisterRoutes(app *fiber.App) {
	app.Post("/user/add", AuthRequired, u.controller.Add)
	app.Get("/user/get/:id", AuthRequired, u.controller.Get)
	app.Put("/user/edit/:id", AuthRequired, u.controller.Edit)
	app.Delete("/user/delete/:id", AuthRequired, u.controller.Delete)
	app.Get("/user/list", AuthRequired, RoleRequired([]string{"all"}), u.controller.List)
	app.Post("/user/upload", AuthRequired, u.controller.SetFile)
}
