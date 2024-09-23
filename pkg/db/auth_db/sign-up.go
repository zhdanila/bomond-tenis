package auth_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type signUpHandler struct {
	pool *sqlx.DB
}

func NewSignUpHandler(pool *sqlx.DB) *signUpHandler {
	return &signUpHandler{pool: pool}
}

func (h *signUpHandler) Exec(ctx context.Context, args *query.SignUpQuery) (err error) {
	query := fmt.Sprintf(`INSERT INTO %s (email, password, username)
		VALUES (:email, :password, :username)
		RETURNING id`, utils.UsersTable)

	rows, err := h.pool.NamedQueryContext(ctx, query, args)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (h *signUpHandler) Context() interface{} {
	return (*query.SignUpQuery)(nil)
}
