package helloRoute

import (
	"io"

	"github.com/go-playground/validator/v10"
)

type PostHelloSchema struct {
	Name string `json:"name" validate:"required,min=3,max=10"`
	Age  int    `json:"age"`
}

func Validate(io.ReadCloser) error {
	var err error
	var payload PostHelloSchema
	validate := validator.New()
	err = validate.Struct(payload)
	return err
}

func PayloadType() PostHelloSchema {
	var x PostHelloSchema
	return x
}
