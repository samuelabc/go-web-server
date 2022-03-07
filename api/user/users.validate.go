package userController

import (
	"github.com/gofrs/uuid"
)

type RegisterUserRequest struct {
	Name     *string `json:"title,omitempty" validate:"required"`
	Password *string `json:"content,omitempty" validate:"required"`
}

type LoginUserRequest struct {
	Name     *string `json:"title,omitempty" validate:"required"`
	Password *string `json:"content,omitempty" validate:"required"`
}

type GetUserRequest struct {
	ID *uuid.UUID `json:"id" validate:"required"`
}

type ListUserRequest struct {
	Name *string `json:"title,omitempty" validate:"required"`
}

type UpdateUserRequest struct {
	ID       *uuid.UUID `json:"id,omitempty" validate:"required"`
	Name     *string    `json:"title,omitempty" validate:"required"`
	Password *string    `json:"content,omitempty" validate:"required"`
}

type DeleteUserRequest struct {
	ID *uuid.UUID `json:"id,omitempty" validate:"required"`
}
