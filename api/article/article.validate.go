package articleController

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	// "gopkg.in/guregu/null.v4"
)

type GetArticleRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type ListArticleRequest struct {
	ID    uuid.UUID      `json:"id,omitempty" validate:""`
	Title sql.NullString `json:"title,omitempty" validate:""`
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
