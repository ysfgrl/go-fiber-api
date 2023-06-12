package response

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/models"
	"path/filepath"
	"runtime"
	"strings"
)

func baseResponse(c *fiber.Ctx, status int, content any, err *models.MyError) error {
	if err != nil {
		return c.Status(status).JSON(models.Response{
			Code:    status,
			Content: content,
			Error:   []*models.MyError{err},
		})
	} else {
		return c.Status(status).JSON(models.Response{
			Code:    status,
			Content: content,
			Error:   []*models.MyError{},
		})
	}

}

func OK(c *fiber.Ctx, content any) error {
	return baseResponse(c, fiber.StatusOK, content, nil)
}

func Unauthorized(c *fiber.Ctx, err *models.MyError) error {
	return baseResponse(c, fiber.StatusUnauthorized, nil, err)
}
func Forbidden(c *fiber.Ctx, err *models.MyError) error {
	return baseResponse(c, fiber.StatusForbidden, nil, err)
}
func NotAllowed(c *fiber.Ctx, content any) error {
	return baseResponse(c, fiber.StatusMethodNotAllowed, content, nil)
}

func NotImplemented(c *fiber.Ctx) error {
	return baseResponse(c, fiber.StatusNotImplemented, nil, nil)
}

func NotFound(c *fiber.Ctx, err *models.MyError) error {
	return baseResponse(c, fiber.StatusNotFound, nil, err)
}

func BadRequest(c *fiber.Ctx, err *models.MyError) error {
	return baseResponse(c, fiber.StatusBadRequest, nil, err)
}

func InternalServerError(c *fiber.Ctx, err *models.MyError) error {
	return baseResponse(c, fiber.StatusInternalServerError, nil, err)
}

func GetError(err error) *models.MyError {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &models.MyError{
		Code:     "code",
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: filepath.Base(function.Name()),
		Line:     line,
		Detail:   err.Error(),
	}
}

func UserError(msg string) *models.MyError {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])
	file, line := function.FileLine(pc[0])
	return &models.MyError{
		Code:     "code",
		File:     strings.Replace(filepath.Base(file), ".go", "", 1),
		Function: function.Name(),
		Line:     line,
		Detail:   msg,
	}
}
