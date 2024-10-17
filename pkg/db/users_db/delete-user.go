package users_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DeleteUserHandler struct {
	pool *sqlx.DB
}

func NewDeleteUserHandler(pool *sqlx.DB) *DeleteUserHandler {
	return &DeleteUserHandler{pool: pool}
}

func (h *DeleteUserHandler) Exec(ctx context.Context, args *query.DeleteUserQuery) (err error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = $1", utils.UsersTable)

	_, err = h.pool.Exec(sql, args.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (h *DeleteUserHandler) Context() interface{} {
	return (*query.DeleteUserQuery)(nil)
}
