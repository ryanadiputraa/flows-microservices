package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(val any) (error, map[string][]string)
}

type validation struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	return &validation{
		validator: validator.New(),
	}
}

func (v *validation) Validate(val any) (error, map[string][]string) {
	err := v.validator.Struct(val)
	errorsMap := make(map[string][]string)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			field := strings.ToLower(fieldErr.Field())
			errorsMap[field] = append(errorsMap[field], FieldErrMsg(fieldErr))

		}
		return errors.New("invalid params"), errorsMap
	}

	return nil, nil
}

func FieldErrMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", strings.ToLower(err.Field()))
	case "max":
		return fmt.Sprintf("%s should have a maximum length of %s", strings.ToLower(err.Field()), err.Param())
	case "min":
		return fmt.Sprintf("%s should have a minimum length of %s", strings.ToLower(err.Field()), err.Param())
	case "email":
		return fmt.Sprintf("%s should be a valid email address", strings.ToLower(err.Field()))
	case "http_url":
		return fmt.Sprintf("%s should be a valid http url", strings.ToLower(err.Field()))
	default:
		return err.Error()
	}
}
