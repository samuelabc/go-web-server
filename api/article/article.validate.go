package articleController

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	// "gopkg.in/guregu/null.v4"
)

type GetArticleRequest struct {
	ID *uuid.UUID `json:"id" validate:"required"`
}

type ListArticleRequest struct {
	ID      *uuid.UUID `json:"id,omitempty" validate:""`
	Title   *string    `json:"title,omitempty" validate:""`
	Content *string    `json:"content,omitempty" validate:""`
}

type UpdateArticleRequest struct {
	ID      *uuid.UUID `json:"id,omitempty" validate:"required"`
	Title   *string    `json:"title,omitempty" validate:""`
	Content *string    `json:"content,omitempty" validate:""`
}

type DeleteArticleRequest struct {
	ID *uuid.UUID `json:"id,omitempty" validate:"required"`
}

type Options struct {
	ID    *uuid.UUID `json:"id,omitempty"`
	Title *string    `json:"title,omitempty"`
}

func parseOptions(jsn []byte) Options {
	var opts Options
	if err := json.Unmarshal(jsn, &opts); err != nil {
		log.Fatal(err)
	}
	// if opts.ID == nil {
	// 	var v uuid.UUID = null
	// 	opts.ID = &v
	// }
	return opts
}
