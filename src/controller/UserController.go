package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/services"
	"go-fiber-api/src/utils/response"
	"go-fiber-api/src/utils/validation"
)

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) Controller {
	return &userController{service: service}
}

func (controller *userController) Add(c *fiber.Ctx) error {

	user := mongo_collections.UserListItem{}
	if err := c.BodyParser(&user); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	if err := validation.Validate(user); err != nil {
		return response.BadRequest(c, err)
	}
	newUser, err := controller.service.AddUser(user)
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
	user, err := controller.service.GetUser(id)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, user)
}

func (controller *userController) List(c *fiber.Ctx) error {

	listRequest := models.ListRequestLastDay()
	if err := c.QueryParser(&listRequest); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	list, err := controller.service.GetList(listRequest)
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

	result, err1 := controller.service.UploadProfile(file)
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
