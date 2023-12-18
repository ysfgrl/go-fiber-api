package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/src/models"
	"go-fiber-api/src/models/elastic_collections"
	"go-fiber-api/src/services"
	"go-fiber-api/src/utils/response"
)

type taskController struct {
	service services.TaskService
}

func NewTaskController(service services.TaskService) Controller {
	return &taskController{
		service: service,
	}
}

func (controller *taskController) Add(c *fiber.Ctx) error {

	taskState := elastic_collections.TaskState{}
	if err := c.BodyParser(&taskState); err != nil {
		return response.BadRequest(c, response.GetError(err))
	}
	newState, err := controller.service.Add(taskState)
	if err != nil {
		return response.NotFound(c, err)
	}
	return response.OK(c, newState)
}

func (controller *taskController) Get(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *taskController) List(c *fiber.Ctx) error {
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

func (controller *taskController) SetFile(c *fiber.Ctx) error {

	return response.NotImplemented(c)
}
func (controller *taskController) Delete(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}

func (controller *taskController) Edit(c *fiber.Ctx) error {
	return response.NotImplemented(c)
}
