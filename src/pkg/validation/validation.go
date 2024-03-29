package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/ysfgrl/go-fiber-api/src/models"
)

var myValidator = validator.New()

func Validate(schema interface{}) *models.Error {
	err := myValidator.Struct(schema)
	if err != nil {
		var errors []models.ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			var el models.ValidationError
			el.Field = err.Field()
			el.Value = err.Value()
			el.Tag = err.Tag()
			el.Param = err.Param()
			errors = append(errors, el)
		}

		return &models.Error{
			Function: "Add",
			File:     "Controller",
			Detail:   errors,
			Code:     "validation.error",
		}
	}
	return nil
}
