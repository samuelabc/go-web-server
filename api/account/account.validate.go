package accountController

import (
	"github.com/gofrs/uuid"
)

type RegisterAccountRequest struct {
	Name     *string `json:"name,omitempty" validate:"required"`
	Password *string `json:"password,omitempty" validate:"required"`
	Email    *string `json:"email,omitempty" validate:""`
}

type LoginAccountRequest struct {
	Name     *string `json:"name,omitempty" validate:"required"`
	Password *string `json:"password,omitempty" validate:"required"`
}

type GetAccountRequest struct {
	ID *uuid.UUID `json:"id" validate:"required"`
}

type ListAccountRequest struct {
	ID    *uuid.UUID `json:"id" validate:""`
	Name  *string    `json:"name,omitempty" validate:""`
	Email *string    `json:"email,omitempty" validate:""`
}

type UpdateAccountRequest struct {
	ID       *uuid.UUID `json:"id,omitempty" validate:"required"`
	Name     *string    `json:"name,omitempty" validate:"required"`
	Password *string    `json:"password,omitempty" validate:"required"`
}

type DeleteAccountRequest struct {
	ID *uuid.UUID `json:"id,omitempty" validate:"required"`
}
