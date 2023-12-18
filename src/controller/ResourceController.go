package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/mongo_collections"
	"go-fiber-api/src/services"
	"go-fiber-api/src/utils/response"
	"go-fiber-api/src/utils/validation"
	"time"
)

type resourceController struct {
	service services.ResourceService
}

func NewResourceController(service services.ResourceService) Controller {
	return &resourceController{service: service}
}

func (controller *resourceController) Add(c *fiber.Ctx) error {

	resource := mongo_collections.ResourceListItem{
		CreatedAt: time.Now().UTC(),
		Download:  false,
		Type:      "undefined",
	}
	if err := c.BodyParser(&resource); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	if err := validation.Validate(resource); err != nil {
		return response.BadRequest(c, err)
	}
	newResource, err := controller.service.AddResource(resource)
	if err != nil {
		return response.InternalServerError(c, err)
	}
	return response.OK(c, newResource)
}

func (controller *resourceController) Get(c *fiber.Ctx) error {
	id := c.Params("id", "id")
	if len(id) != 24 {
		return response.BadRequest(c, response.UserError("id required"))
	}
	resource, err := controller.service.GetResource(id)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, resource)
}

func (controller *resourceController) List(c *fiber.Ctx) error {

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

func (controller *resourceController) SetFile(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	result, err1 := controller.service.UploadResource(file)
	if err1 != nil {
		return response.InternalServerError(c, err1)
	}
	return response.OK(c, result)
}
func (controller *resourceController) Delete(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *resourceController) Edit(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}
