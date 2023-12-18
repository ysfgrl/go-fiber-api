package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/services"
	"go-fiber-api/src/utils/response"
	"go-fiber-api/src/utils/security"
	"go-fiber-api/src/utils/validation"
)

type authController struct {
	service services.UserService
}

func NewAuthController(service services.UserService) Controller {
	return &authController{service: service}
}

func (controller *authController) Get(c *fiber.Ctx) error {

	authModel := models.SignIn{}

	if "POST" == c.Method() {
		if err2 := c.BodyParser(&authModel); err2 != nil {
			return response.BadRequest(c, response.GetError(err2))
		}
	} else if "GET" == c.Method() {
		if err1 := c.QueryParser(&authModel); err1 != nil {
			return response.BadRequest(c, response.GetError(err1))
		}
	}

	if err := validation.Validate(authModel); err != nil {
		return response.BadRequest(c, err)
	}

	user, err := controller.service.GetByUserName(authModel.Username)
	if err != nil {
		return response.InternalServerError(c, err)
	}

	err = security.VerifyPassword(user.Password, authModel.Password)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	token, err := security.NewToken(*user, "secret")
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, models.TokenModel{Token: token})
}
func (controller *authController) Add(c *fiber.Ctx) error {
	signUpModel := mongo_collections.UserListItem{}
	if err := c.BodyParser(&signUpModel); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	if err := validation.Validate(signUpModel); err != nil {
		return response.BadRequest(c, err)
	}

	existUser, err := controller.service.GetByUserName(signUpModel.UserName)
	if existUser != nil {
		return response.NotFound(c, err)
	}
	signUpModel.Password, err = security.EncryptPassword(signUpModel.Password)
	if err != nil {
		return response.BadRequest(c, err)
	}
	newUser, err := controller.service.AddUser(signUpModel)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, newUser)
}

func (controller *authController) Delete(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *authController) Edit(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *authController) List(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *authController) SetFile(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}
