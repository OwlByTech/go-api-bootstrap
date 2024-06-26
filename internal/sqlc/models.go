// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
	"time"
)

type Client struct {
	ClientID  int64        `json:"client_id"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	GivenName string       `json:"given_name"`
	Surname   string       `json:"surname"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
