package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/controller"
)

type userRouter struct {
	controller controller.Controller
}

func NewUserRoute(baseController controller.Controller) Route {
	return &userRouter{controller: baseController}
}

func (u userRouter) RegisterRoutes(app *fiber.App) {
	app.Post("/user/add", AuthRequired, u.controller.Add)
	app.Get("/user/get/:id", u.controller.Get)
	app.Put("/user/edit/:id", u.controller.Edit)
	app.Delete("/user/delete/:id", u.controller.Delete)
	app.Get("/user/list", RoleRequired([]string{"all"}), u.controller.List)
	app.Post("/user/upload", AuthRequired, u.controller.SetFile)
}

func (u *userRouter) GetController() *controller.Controller {
	return &u.controller
}
