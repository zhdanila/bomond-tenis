package users_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UpdateUserHandler struct {
	pool *sqlx.DB
}

func NewUpdateUserHandler(pool *sqlx.DB) *UpdateUserHandler {
	return &UpdateUserHandler{pool: pool}
}

func (h *UpdateUserHandler) Exec(ctx context.Context, args *query.UpdateUserQuery) (err error) {
	sql := fmt.Sprintf("UPDATE %s SET email=$1, password=$2, username=$3 WHERE id=$4", utils.UsersTable)

	_, err = h.pool.Exec(sql, args.Email, args.Password, args.Username, args.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (h *UpdateUserHandler) Context() interface{} {
	return (*query.UpdateUserQuery)(nil)
}
