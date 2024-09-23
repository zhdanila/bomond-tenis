package auth_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type signInHandler struct {
	pool *sqlx.DB
}

func NewSignInHandler(pool *sqlx.DB) *signInHandler {
	return &signInHandler{pool: pool}
}

func (h *signInHandler) Exec(ctx context.Context, args *query.SignInQuery) (err error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE email = :email", utils.UsersTable)

	rows, err := h.pool.NamedQueryContext(ctx, sql, args)
	if err != nil {
		return err
	}

	defer rows.Close()

	var a query.Account
	for rows.Next() {
		if err = rows.Scan(&a.Email, &a.Username, &a.Password, &a.ID); err != nil {
			return err
		}

		args.Out.Account = a
	}

	return nil
}

func (h *signInHandler) Context() interface{} {
	return (*query.SignInQuery)(nil)
}
