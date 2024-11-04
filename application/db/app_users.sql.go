// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: app_users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const CountUsersByOrg = `-- name: CountUsersByOrg :one
SELECT COUNT(*) as total 
FROM app_users u
WHERE u.organization_id = $1
AND (u.display_name ILIKE '%' || $2 || '%' OR u.email ILIKE '%' || $2 || '%' OR u.username ILIKE '%' || $2 || '%')
`

type CountUsersByOrgParams struct {
	OrganizationID int32       `db:"organization_id" json:"organization_id"`
	Column2        pgtype.Text `db:"column_2" json:"column_2"`
}

func (q *Queries) CountUsersByOrg(ctx context.Context, arg *CountUsersByOrgParams) (int64, error) {
	row := q.db.QueryRow(ctx, CountUsersByOrg, arg.OrganizationID, arg.Column2)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const CreateUser = `-- name: CreateUser :one
INSERT INTO app_users (
  organization_id,
  username,
  display_name,
  email,
  password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING user_id, organization_id, username, display_name, email, password, created_at, updated_at, is_active, is_locked
`

type CreateUserParams struct {
	OrganizationID int32  `db:"organization_id" json:"organization_id"`
	Username       string `db:"username" json:"username"`
	DisplayName    string `db:"display_name" json:"display_name"`
	Email          string `db:"email" json:"email"`
	Password       string `db:"password" json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (*AppUser, error) {
	row := q.db.QueryRow(ctx, CreateUser,
		arg.OrganizationID,
		arg.Username,
		arg.DisplayName,
		arg.Email,
		arg.Password,
	)
	var i AppUser
	err := row.Scan(
		&i.UserID,
		&i.OrganizationID,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsActive,
		&i.IsLocked,
	)
	return &i, err
}

const DeleteUser = `-- name: DeleteUser :exec
DELETE FROM app_users
WHERE user_id = $1 AND organization_id = $2
`

type DeleteUserParams struct {
	UserID         int32 `db:"user_id" json:"user_id"`
	OrganizationID int32 `db:"organization_id" json:"organization_id"`
}

func (q *Queries) DeleteUser(ctx context.Context, arg *DeleteUserParams) error {
	_, err := q.db.Exec(ctx, DeleteUser, arg.UserID, arg.OrganizationID)
	return err
}

const GetUserByID = `-- name: GetUserByID :one
SELECT u.user_id, u.organization_id, u.username, u.display_name, u.email, u.password, u.created_at, u.updated_at, u.is_active, u.is_locked
FROM app_users u
WHERE u.user_id = $1
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID int32) (*AppUser, error) {
	row := q.db.QueryRow(ctx, GetUserByID, userID)
	var i AppUser
	err := row.Scan(
		&i.UserID,
		&i.OrganizationID,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsActive,
		&i.IsLocked,
	)
	return &i, err
}

const GetUserByName = `-- name: GetUserByName :one
SELECT u.user_id, u.organization_id, u.username, u.display_name, u.email, u.password, u.created_at, u.updated_at, u.is_active, u.is_locked 
FROM app_users u
WHERE u.username = $1
LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (*AppUser, error) {
	row := q.db.QueryRow(ctx, GetUserByName, username)
	var i AppUser
	err := row.Scan(
		&i.UserID,
		&i.OrganizationID,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsActive,
		&i.IsLocked,
	)
	return &i, err
}

const GetUserByOrg = `-- name: GetUserByOrg :one
SELECT u.user_id, u.organization_id, u.username, u.display_name, u.email, u.password, u.created_at, u.updated_at, u.is_active, u.is_locked
FROM app_users u
WHERE u.user_id = $1 AND u.organization_id = $2
LIMIT 1
`

type GetUserByOrgParams struct {
	UserID         int32 `db:"user_id" json:"user_id"`
	OrganizationID int32 `db:"organization_id" json:"organization_id"`
}

func (q *Queries) GetUserByOrg(ctx context.Context, arg *GetUserByOrgParams) (*AppUser, error) {
	row := q.db.QueryRow(ctx, GetUserByOrg, arg.UserID, arg.OrganizationID)
	var i AppUser
	err := row.Scan(
		&i.UserID,
		&i.OrganizationID,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsActive,
		&i.IsLocked,
	)
	return &i, err
}

const ListUsersByOrg = `-- name: ListUsersByOrg :many
SELECT u.user_id, u.organization_id, u.username, u.display_name, u.email, u.password, u.created_at, u.updated_at, u.is_active, u.is_locked 
FROM app_users u
WHERE u.organization_id = $1
AND (u.display_name ILIKE '%' || $2 || '%' OR u.email ILIKE '%' || $2 || '%' OR u.username ILIKE '%' || $2 || '%')
ORDER BY u.display_name
LIMIT $4 OFFSET $3
`

type ListUsersByOrgParams struct {
	OrganizationID int32       `db:"organization_id" json:"organization_id"`
	Column2        pgtype.Text `db:"column_2" json:"column_2"`
	Offset         int32       `db:"offset" json:"offset"`
	Limit          int32       `db:"limit" json:"limit"`
}

func (q *Queries) ListUsersByOrg(ctx context.Context, arg *ListUsersByOrgParams) ([]*AppUser, error) {
	rows, err := q.db.Query(ctx, ListUsersByOrg,
		arg.OrganizationID,
		arg.Column2,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*AppUser{}
	for rows.Next() {
		var i AppUser
		if err := rows.Scan(
			&i.UserID,
			&i.OrganizationID,
			&i.Username,
			&i.DisplayName,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsActive,
			&i.IsLocked,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateUser = `-- name: UpdateUser :exec
UPDATE app_users
SET
  display_name = $1,
  email = $2,
  updated_at = NOW()
WHERE user_id = $3 AND organization_id = $4
RETURNING user_id, organization_id, username, display_name, email, password, created_at, updated_at, is_active, is_locked
`

type UpdateUserParams struct {
	DisplayName    string `db:"display_name" json:"display_name"`
	Email          string `db:"email" json:"email"`
	UserID         int32  `db:"user_id" json:"user_id"`
	OrganizationID int32  `db:"organization_id" json:"organization_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg *UpdateUserParams) error {
	_, err := q.db.Exec(ctx, UpdateUser,
		arg.DisplayName,
		arg.Email,
		arg.UserID,
		arg.OrganizationID,
	)
	return err
}
