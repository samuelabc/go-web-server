package authModel

import "time"

type Token struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	AccountID int       `json:"-"`

	Token      string    `json:"-"`
	Expiry     time.Time `json:"-"`
	Mobile     bool      `sql:",notnull" json:"mobile"`
	Identifier string    `json:"identifier,omitempty"`
}
