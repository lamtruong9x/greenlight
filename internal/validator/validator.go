package validator

import (
	"errors"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lamtruong9x/greenlight/internal/data"
)

var validate *validator.Validate

func New() (*validator.Validate, error) {
	validate = validator.New()
	if err := validate.RegisterValidation("lteyear", LessThanCurrentYear); err != nil {
		return nil, errors.New("cannot register custom validations (should be paniced)")
	}
	validate.RegisterCustomTypeFunc(runtimeCustomTypeFunc, data.Runtime(1))
	return validate, nil
}

func LessThanCurrentYear(fl validator.FieldLevel) bool {
	return fl.Field().Int() <= int64(time.Now().Year())
}

func runtimeCustomTypeFunc(field reflect.Value) any {
	if value, ok := field.Interface().(data.Runtime); ok {
		return int32(value)
	}
	return nil
}
