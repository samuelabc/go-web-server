// Package models contains application specific entities.
package models

import (
	"time"

	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"

	"github.com/go-pg/pg/orm"
	"github.com/go-playground/validator"
)

type ID string
type Title string
type Content string
type CreatedAt time.Time
type UpdatedAt time.Time

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// BeforeInsert hook executed before database insert operation.
func (p *Article) BeforeInsert(db orm.DB) *errorModel.AppError {
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook executed before database update operation.
func (p *Article) BeforeUpdate(db orm.DB) *errorModel.AppError {
	p.UpdatedAt = time.Now()
	return p.Validate()
}

// Validate validates Article struct and returns validation errors.
func (p *Article) Validate() *errorModel.AppError {
	validate := validator.New()
	err := validate.Struct(p)
	return errorHelper.ErrValidation(err)
}
