package schemas

import "github.com/go-playground/validator"

type (
	Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ErrorValidate struct {
		FailedField string `json:"failed_field"`
		Tag         string `json:"tag"`
		Value       string `json:"value"`
	}
)

var validate = validator.New()

func ValidateStruct(data interface{}) []*ErrorValidate {
	var errors []*ErrorValidate

	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationErr := range validationErrors {
			errorItem := &ErrorValidate{
				FailedField: validationErr.StructNamespace(),
				Tag:         validationErr.Tag(),
				Value:       validationErr.Param(),
			}
			errors = append(errors, errorItem)
		}
	}

	return errors
}
