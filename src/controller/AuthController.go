package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/pkg/hash"
	"github.com/ysfgrl/go-fiber-api/src/pkg/response"
	"github.com/ysfgrl/go-fiber-api/src/pkg/validation"
	"github.com/ysfgrl/go-fiber-api/src/repository"
)

type authController struct {
}

func NewAuthController() *authController {
	return &authController{}
}

func (controller *authController) get(c *fiber.Ctx) error {

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

	user, err := repository.UserRepo.GetByFirst(c.UserContext(), "userName", authModel.Username)
	if err != nil {
		return response.InternalServerError(c, err)
	}

	err = hash.VerifyPassword(user.Password, authModel.Password)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	token, err := hash.NewToken(*user, "secret")
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, models.TokenModel{Token: token})
}

func (controller *authController) add(c *fiber.Ctx) error {
	//signUpModel := user_repository.UserListItem{}
	//if err := c.BodyParser(&signUpModel); err != nil {
	//	return response.BadRequest(c, response.GetError(err))
	//}
	//if err := validation.Validate(signUpModel); err != nil {
	//	return response.BadRequest(c, err)
	//}
	//
	//existUser, err := repository.UserRepo.GetByFirst(c.UserContext(),"userName",signUpModel.UserName)
	//if existUser != nil {
	//	return response.NotFound(c, err)
	//}
	//signUpModel.Password, err = security.EncryptPassword(signUpModel.Password)
	//if err != nil {
	//	return response.BadRequest(c, err)
	//}
	//newUser, err := controller.service.AddUser(signUpModel)
	//if err != nil {
	//	return response.InternalServerError(c, err)
	//}
	//return response.OK(c, newUser)
	return response.NotImplemented(c)
}

func (controller *authController) delete(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *authController) edit(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *authController) list(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *authController) setFile(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}
