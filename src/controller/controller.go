package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/go-fiber-api/src/middleware"
)

var (
	UserController *userController = nil
	AuthController *authController = nil
)

func init() {
	UserController = NewUserController()
	AuthController = NewAuthController()
}

func (controller *userController) Routes(app *fiber.App) {
	app.Post("/user/add", middleware.AuthRequired, controller.addUser)
	app.Get("/user/get/:id", controller.getUser)
	app.Put("/user/edit/:id", controller.editUser)
	app.Delete("/user/delete/:id", controller.deleteUser)
	app.Get("/user/list", middleware.RoleRequired([]string{"all"}), controller.listUser)
	app.Post("/user/upload", middleware.AuthRequired, controller.setFile)
}

func (controller *authController) Routes(app *fiber.App) {
	app.Post("/auth/signup", controller.add)
	app.Get("/auth/signup", controller.add)
	app.Add("POST", "/auth/signin", controller.get)
	app.Add("GET", "/auth/signin", controller.get)
	app.Delete("/auth/signout", controller.delete)
}
