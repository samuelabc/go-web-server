package accountController

import (
	"github.com/gofrs/uuid"
)

type RegisterAccountRequest struct {
	Name     *string `json:"title,omitempty" validate:"required"`
	Password *string `json:"content,omitempty" validate:"required"`
}

type LoginAccountRequest struct {
	Name     *string `json:"title,omitempty" validate:"required"`
	Password *string `json:"content,omitempty" validate:"required"`
}

type GetAccountRequest struct {
	ID *uuid.UUID `json:"id" validate:"required"`
}

type ListAccountRequest struct {
	Name *string `json:"title,omitempty" validate:"required"`
}

type UpdateAccountRequest struct {
	ID       *uuid.UUID `json:"id,omitempty" validate:"required"`
	Name     *string    `json:"title,omitempty" validate:"required"`
	Password *string    `json:"content,omitempty" validate:"required"`
}

type DeleteAccountRequest struct {
	ID *uuid.UUID `json:"id,omitempty" validate:"required"`
}
