package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/models"
	"go-fiber-api/myUtils/response"
	"go-fiber-api/myUtils/validation"
	"go-fiber-api/services"
)

type userController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) BaseController {
	return &userController{userService: service}
}

func (controller *userController) Add(c *fiber.Ctx) error {

	user := models.UserListItem{}
	if err := c.BodyParser(&user); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	if err := validation.Validate(user); err != nil {
		return response.BadRequest(c, err)
	}
	newUser, err := controller.userService.AddUser(user)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, newUser)
}

func (controller *userController) Get(c *fiber.Ctx) error {

	id := c.Params("id", "id")
	if len(id) != 24 {
		return response.BadRequest(c, response.UserError("id required"))
	}
	user, err := controller.userService.GetUser(id)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, user)
}

func (controller *userController) List(c *fiber.Ctx) error {

	listRequest := models.DefaultListResponse()
	if err := c.QueryParser(&listRequest); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	list, err := controller.userService.GetList(listRequest)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, list)
}

func (controller *userController) SetFile(c *fiber.Ctx) error {
	file, err := c.FormFile("profile")
	if err != nil {
		return response.BadRequest(c, response.GetError(err))
	}

	result, err1 := controller.userService.UploadProfile(file)
	if err1 != nil {
		return response.InternalServerError(c, err1)
	}
	return response.OK(c, result)
}
func (controller *userController) Delete(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *userController) Edit(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}
