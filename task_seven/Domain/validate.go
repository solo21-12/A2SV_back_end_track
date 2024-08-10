package domain

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (t *User) Validate() error {
	return validate.Struct(t)
}

func (t *UserCreateRequest) Validate() error {
	return validate.Struct(t)
}

func (t *TaskCreateDTO) Validate() error {
	return validate.Struct(t)
}
