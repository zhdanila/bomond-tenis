package auth_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SignUpHandler struct {
	pool *sqlx.DB
}

func NewSignUpHandler(pool *sqlx.DB) *SignUpHandler {
	return &SignUpHandler{pool: pool}
}

func (h *SignUpHandler) Exec(ctx context.Context, args *query.SignUpQuery) (err error) {
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

func (h *SignUpHandler) Context() interface{} {
	return (*query.SignUpQuery)(nil)
}
