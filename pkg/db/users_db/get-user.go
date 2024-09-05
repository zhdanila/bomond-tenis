package users_db

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type getUserHandler struct {
	pool *sqlx.DB
}

func NewGetUserHandler(pool *sqlx.DB) *getUserHandler {
	return &getUserHandler{pool: pool}
}

func (h *getUserHandler) Exec(ctx context.Context, args *query.GetUserQuery) (err error) {
	fmt.Println("get user")

	return nil
}

func (h *getUserHandler) Context() interface{} {
	return (*query.GetUserQuery)(nil)
}
