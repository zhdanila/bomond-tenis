package users_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GetUserHandler struct {
	pool *sqlx.DB
}

func NewGetUserHandler(pool *sqlx.DB) *GetUserHandler {
	return &GetUserHandler{pool: pool}
}

func (h *GetUserHandler) Exec(ctx context.Context, args *query.GetUserQuery) (err error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", utils.UsersTable)

	rows := h.pool.QueryRow(sql, args.UserId)

	var account query.Account
	if err = rows.Scan(&account.ID, &account.Email, &account.Password, &account.Username); err != nil {
		return err
	}

	args.Out.Account = account

	return nil
}

func (h *GetUserHandler) Context() interface{} {
	return (*query.GetUserQuery)(nil)
}
