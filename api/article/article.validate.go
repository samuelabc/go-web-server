package articleController

import (
	"github.com/gofrs/uuid"
)

type GetArticleRequest struct {
	ID *uuid.UUID `json:"id" validate:"required"`
}

type ListArticleRequest struct {
	ID      *uuid.UUID `json:"id,omitempty" validate:""`
	Title   *string    `json:"title,omitempty" validate:""`
	Content *string    `json:"content,omitempty" validate:""`
}

type CreateArticleRequest struct {
	Title   *string `json:"title,omitempty" validate:"required"`
	Content *string `json:"content,omitempty" validate:""`
}

type UpdateArticleRequest struct {
	ID      *uuid.UUID `json:"id,omitempty" validate:"required"`
	Title   *string    `json:"title,omitempty" validate:""`
	Content *string    `json:"content,omitempty" validate:""`
}

type DeleteArticleRequest struct {
	ID *uuid.UUID `json:"id,omitempty" validate:"required"`
}
