package auth_db

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type logoutHandler struct {
	pool *sqlx.DB
}

func NewLogoutHandler(pool *sqlx.DB) *logoutHandler {
	return &logoutHandler{pool: pool}
}

func (h *logoutHandler) Exec(ctx context.Context, args *query.LogoutQuery) (err error) {
	fmt.Println("logout")

	return nil
}

func (h *logoutHandler) Context() interface{} {
	return (*query.LogoutQuery)(nil)
}

//query := fmt.Sprintf(`INSERT INTO %s (uuid, account, type, parameters, data, status, created_at, send_at)
//	VALUES (gen_random_uuid(), :account, :type, :parameters, :data, :status, :created_at, :send_at)
//	RETURNING uuid`, NotificationTable)
//
//rows, err := h.pool.NamedQueryContext(ctx, query, params)
//if err != nil {
//	return err
//}
//
//defer rows.Close()
//
//for rows.Next() {
//	var u string
//
//	if err = rows.Scan(&u); err != nil {
//		return err
//	}
//
//	uu, err := uuid.Parse(u)
//	if err != nil {
//		return err
//	}
//
//	args.Uuid = uu
//}
