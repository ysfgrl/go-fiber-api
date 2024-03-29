package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysfgrl/go-fiber-api/src/models"
)

func baseResponse(c *fiber.Ctx, status int, content any, err *models.Error) error {
	if err != nil {
		return c.Status(status).JSON(models.Response{
			Code:    status,
			Content: content,
			Error:   []*models.Error{err},
		})
	} else {
		return c.Status(status).JSON(models.Response{
			Code:    status,
			Content: content,
			Error:   []*models.Error{},
		})
	}

}

func OK(c *fiber.Ctx, content any) error {
	return baseResponse(c, fiber.StatusOK, content, nil)
}

func Unauthorized(c *fiber.Ctx, err *models.Error) error {
	return baseResponse(c, fiber.StatusUnauthorized, nil, err)
}
func Forbidden(c *fiber.Ctx, err *models.Error) error {
	return baseResponse(c, fiber.StatusForbidden, nil, err)
}
func NotAllowed(c *fiber.Ctx, content any) error {
	return baseResponse(c, fiber.StatusMethodNotAllowed, content, nil)
}

func NotImplemented(c *fiber.Ctx) error {
	return baseResponse(c, fiber.StatusNotImplemented, nil, models.UserError("NotImplemented"))
}

func NotFound(c *fiber.Ctx, err *models.Error) error {
	return baseResponse(c, fiber.StatusNotFound, nil, err)
}

func BadRequest(c *fiber.Ctx, err *models.Error) error {
	return baseResponse(c, fiber.StatusBadRequest, nil, err)
}

func InternalServerError(c *fiber.Ctx, err *models.Error) error {
	return baseResponse(c, fiber.StatusInternalServerError, nil, err)
}
