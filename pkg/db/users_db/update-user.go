package users_db

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type updateUserHandler struct {
	pool *sqlx.DB
}

func NewUpdateUserHandler(pool *sqlx.DB) *updateUserHandler {
	return &updateUserHandler{pool: pool}
}

func (h *updateUserHandler) Exec(ctx context.Context, args *query.UpdateUserQuery) (err error) {
	fmt.Println("update user")

	return nil
}

func (h *updateUserHandler) Context() interface{} {
	return (*query.UpdateUserQuery)(nil)
}
