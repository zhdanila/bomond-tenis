package auth_db

import (
	db2 "bomond-tenis/pkg/db"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	NotificationTable = "notification"
)

type logoutHandler struct {
	pool *sqlx.DB
}

func NewLogoutHandler(pool *sqlx.DB) *logoutHandler {
	return &logoutHandler{pool: pool}
}

func (h *logoutHandler) Exec(ctx context.Context, args *db2.LogoutQuery) (err error) {
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

	fmt.Println("logout")

	return nil
}

func (h *logoutHandler) Context() interface{} {
	return (*db2.LogoutQuery)(nil)
}
