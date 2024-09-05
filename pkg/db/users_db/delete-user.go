package users_db

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type deleteUserHandler struct {
	pool *sqlx.DB
}

func NewDeleteUserHandler(pool *sqlx.DB) *deleteUserHandler {
	return &deleteUserHandler{pool: pool}
}

func (h *deleteUserHandler) Exec(ctx context.Context, args *query.DeleteUserQuery) (err error) {
	fmt.Println("delete user")

	return nil
}

func (h *deleteUserHandler) Context() interface{} {
	return (*query.DeleteUserQuery)(nil)
}
