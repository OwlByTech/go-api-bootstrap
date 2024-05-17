// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: client.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createClient = `-- name: CreateClient :one
INSERT INTO client (email, password, given_name, surname, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING client_id, email, password, given_name, surname, created_at, updated_at, deleted_at
`

type CreateClientParams struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	GivenName string    `json:"given_name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.Email,
		arg.Password,
		arg.GivenName,
		arg.Surname,
		arg.CreatedAt,
	)
	var i Client
	err := row.Scan(
		&i.ClientID,
		&i.Email,
		&i.Password,
		&i.GivenName,
		&i.Surname,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAllClients = `-- name: DeleteAllClients :execresult
DELETE FROM client
`

func (q *Queries) DeleteAllClients(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteAllClients)
}

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM client
WHERE client_id = $1
`

func (q *Queries) DeleteClient(ctx context.Context, clientID int64) error {
	_, err := q.db.ExecContext(ctx, deleteClient, clientID)
	return err
}

const getClient = `-- name: GetClient :one
SELECT client_id, email, password, given_name, surname, created_at, updated_at, deleted_at FROM client
WHERE client_id = $1 LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, clientID int64) (Client, error) {
	row := q.db.QueryRowContext(ctx, getClient, clientID)
	var i Client
	err := row.Scan(
		&i.ClientID,
		&i.Email,
		&i.Password,
		&i.GivenName,
		&i.Surname,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getClientByEmail = `-- name: GetClientByEmail :one
SELECT email, password, given_name, surname
FROM client
WHERE email = $1
`

type GetClientByEmailRow struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
}

func (q *Queries) GetClientByEmail(ctx context.Context, email string) (GetClientByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getClientByEmail, email)
	var i GetClientByEmailRow
	err := row.Scan(
		&i.Email,
		&i.Password,
		&i.GivenName,
		&i.Surname,
	)
	return i, err
}

const listClients = `-- name: ListClients :many
SELECT client_id, email, password, given_name, surname, created_at, updated_at, deleted_at FROM client
ORDER BY given_name
`

func (q *Queries) ListClients(ctx context.Context) ([]Client, error) {
	rows, err := q.db.QueryContext(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ClientID,
			&i.Email,
			&i.Password,
			&i.GivenName,
			&i.Surname,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClientById = `-- name: UpdateClientById :exec
UPDATE client
SET email = $2, password = $3, given_name = $4, surname = $5, updated_at = $6
WHERE client_id = $1
`

type UpdateClientByIdParams struct {
	ClientID  int64        `json:"client_id"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	GivenName string       `json:"given_name"`
	Surname   string       `json:"surname"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func (q *Queries) UpdateClientById(ctx context.Context, arg UpdateClientByIdParams) error {
	_, err := q.db.ExecContext(ctx, updateClientById,
		arg.ClientID,
		arg.Email,
		arg.Password,
		arg.GivenName,
		arg.Surname,
		arg.UpdatedAt,
	)
	return err
}