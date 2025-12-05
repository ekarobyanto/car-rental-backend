package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func injectFilesFromRequest[T any](c *fiber.Ctx, req *T) {
	rv := reflect.ValueOf(req).Elem()
	rt := reflect.TypeOf(req).Elem()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		formTag := field.Tag.Get("form")

		if field.Type == reflect.TypeOf((*multipart.FileHeader)(nil)) && formTag != "" {
			if file, err := c.FormFile(formTag); err == nil && file != nil && file.Filename != "" {
				rv.Field(i).Set(reflect.ValueOf(file))
			} else {
				rv.Field(i).Set(reflect.Zero(field.Type))
			}
		}
	}
}

func ParseValidateRequest[T any](c *fiber.Ctx) (T, error) {
	var req T
	err := c.BodyParser(&req)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError
		var invalidUnmarshalErr *json.InvalidUnmarshalError
		var unsupportedTypeErr *json.UnsupportedTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return req, fmt.Errorf("malformed JSON at position %d", syntaxErr.Offset)

		case errors.As(err, &unmarshalTypeErr):
			return req, fmt.Errorf("invalid type for field '%s': expected %s got %s",
				unmarshalTypeErr.Field,
				unmarshalTypeErr.Type.String(),
				unmarshalTypeErr.Value,
			)

		case errors.As(err, &invalidUnmarshalErr):
			return req, fmt.Errorf("cannot unmarshal JSON into the provided type")

		case errors.As(err, &unsupportedTypeErr):
			return req, fmt.Errorf("unsupported type for JSON unmarshaling: %s", unsupportedTypeErr.Type.String())

		case strings.Contains(err.Error(), "EOF"):
			return req, fmt.Errorf("empty JSON body")

		case strings.Contains(err.Error(), "cannot unmarshal"):
			return req, fmt.Errorf("invalid JSON structure: %v", err.Error())

		default:
			return req, fmt.Errorf("invalid input: %v", err.Error())
		}
	}

	injectFilesFromRequest(c, &req)

	err = validate.Struct(req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var errMsgs []string
			for _, fe := range ve {
				errMsg := fmt.Sprintf("field '%s' failed validation tag '%s'", fe.Field(), fe.Tag())
				errMsgs = append(errMsgs, errMsg)
			}
			return req, errors.New(strings.Join(errMsgs, "; "))
		}
		return req, err
	}

	return req, nil
}
