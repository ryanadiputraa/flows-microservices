package validator

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(val any) (error, map[string][]string)
}

type validation struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	v := validator.New()
	_ = v.RegisterValidation("ISO8601date", isISO8601Date)

	return &validation{
		validator: v,
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
		return err, errorsMap
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
	case "ISO8601date":
		return fmt.Sprintf("%s should be a valid ISO8601 date format", strings.ToLower(err.Field()))
	default:
		return err.Error()
	}
}

func isISO8601Date(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339Nano, fl.Field().String())
	return err == nil
}
