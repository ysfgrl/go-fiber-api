package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/go-fiber-api/src/models"
	"github.com/ysfgrl/go-fiber-api/src/pkg/hash"
	"github.com/ysfgrl/go-fiber-api/src/pkg/response"
	"github.com/ysfgrl/go-fiber-api/src/pkg/validation"
	"github.com/ysfgrl/go-fiber-api/src/repository"
	"github.com/ysfgrl/go-fiber-api/src/repository/user_repository"
	"time"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (controller *userController) addUserBasic(c *fiber.Ctx) error {

	user := models.UserAddBasic{}
	if err := c.BodyParser(&user); err != nil {
		return response.BadRequest(c, models.GetError(err))
	}
	if err := validation.Validate(user); err != nil {
		return response.BadRequest(c, err)
	}

	hp, err := hash.EncryptPassword(user.Password)
	if err := validation.Validate(user); err != nil {
		return response.InternalServerError(c, err)
	}
	newUser := user_repository.User{
		UserName:  "",
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  hp,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	savedUser, err := repository.UserRepo.Add(c.UserContext(), newUser)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, savedUser)
}

func (controller *userController) addUser(c *fiber.Ctx) error {

	user := user_repository.User{}
	if err := c.BodyParser(&user); err != nil {
		return response.BadRequest(c, models.GetError(err))
	}
	if err := validation.Validate(user); err != nil {
		return response.BadRequest(c, err)
	}
	newUser, err := repository.UserRepo.Add(c.UserContext(), user)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, newUser)
}

func (controller *userController) getUser(c *fiber.Ctx) error {

	id := c.Params("id", "id")
	if len(id) != 24 {
		return response.BadRequest(c, models.UserError("id required"))
	}
	user, err := repository.UserRepo.Get(c.UserContext(), id)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, user)
}

func (controller *userController) listUser(c *fiber.Ctx) error {

	listRequest := models.ListRequestLastDay()
	if err := c.QueryParser(&listRequest); err != nil {
		return response.BadRequest(c, models.GetError(err))
	}
	list, err := repository.UserRepo.List(c.UserContext(), listRequest)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, list)
}

func (controller *userController) setFile(c *fiber.Ctx) error {
	//file, err := c.FormFile("profile")
	//if err != nil {
	//	return response.BadRequest(c, response.GetError(err))
	//}
	//
	//result, err1 := repository.UserRepo.UploadProfile(file)
	//if err1 != nil {
	//	return response.InternalServerError(c, err1)
	//}
	//return response.OK(c, result)
	return response.NotImplemented(c)
}
func (controller *userController) deleteUser(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *userController) editUser(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}
