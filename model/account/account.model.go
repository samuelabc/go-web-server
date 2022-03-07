// Package models contains application specific entities.
package accountModel

import (
	"time"

	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"

	"github.com/go-pg/pg/orm"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
)

type Account struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"title"`
	PasswordHash string    `json:"content"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Account) BeforeInsert(db orm.DB) *errorModel.AppError {
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Account) BeforeUpdate(db orm.DB) *errorModel.AppError {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Article struct and returns validation errors.
func (p *Account) Validate() *errorModel.AppError {
	validate := validator.New()
	err := validate.Struct(p)
	return errorHelper.ErrValidation(err)
}
