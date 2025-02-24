package validations

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

func GetMsgErr(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Email is invalid"
	case "min":
		return "Should be less than " + fe.Param()
	case "max":
		return "Should be greater than " + fe.Param()
	}

	return "Unknown error"
}

func ValidateRequest(err error, data interface{}) []ErrMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrMsg, len(ve))
		for i, fe := range ve {
			field, _ := reflect.TypeOf(data).FieldByName(fe.Field())
			jsonTag := field.Tag.Get("json")
			out[i] = ErrMsg{
				Field:   fe.Field(),
				Message: GetMsgErr(fe),
				Name:    jsonTag,
			}
		}
		return out
	}
	return nil
}
